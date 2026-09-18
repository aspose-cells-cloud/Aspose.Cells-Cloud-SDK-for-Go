#!/usr/bin/env python3
"""Convert the legacy-model integration tests (integrationtests/) into
new-model tests (newtests/).

Authoritative sources:
  * aspose.cells.cloud.specification.json  -- operation parameters (name /
    required / data type)
  * requests/*.go                          -- generated request constructors
    (positional parameter order + option value types)
  * models/*.go                            -- model field types (pointer
    conversion for body models)

Run from the repository root:  python3 tools/convert_integrationtests.py
"""

import json
import os
import re
import subprocess
import sys
import textwrap

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
SPEC_PATH = os.path.join(ROOT, "aspose.cells.cloud.specification.json")
REQ_DIR = os.path.join(ROOT, "requests")
MODEL_DIR = os.path.join(ROOT, "models")
SRC_DIR = os.path.join(ROOT, "integrationtests")
OUT_DIR = os.path.join(ROOT, "newtests")

PKG = "asposecellscloud"

# ---------------------------------------------------------------- helpers


def norm(name):
    """Loose comparison key for parameter / field names."""
    return name.lower().rstrip("_").replace("_", "")


# Legacy request fields that the new model renamed. Keyed by normalized name,
# applied only when the operation actually exposes the new name.
ALIASES = {
    # legacy: the local file to convert; new model reuses the spec's `File`.
    "localpath": "file",
}


# ---------------------------------------------------------------- spec


class Spec(object):
    def __init__(self):
        with open(SPEC_PATH, encoding="utf-8") as fh:
            data = json.load(fh)
        self.ops = {}
        for op in data["Operations"]:
            params = []
            for p in op.get("Parameters") or []:
                params.append(
                    {
                        "name": p["Name"],
                        "required": bool(p.get("Required")),
                        "type": (p.get("DataType") or {}).get("Identifier") or "String",
                        "group": p.get("Group") or "",
                    }
                )
            self.ops[op["Name"]] = {
                "params": params,
                "method": op.get("HttpMethod"),
                "path": op.get("Path"),
            }

    def operation(self, name):
        return self.ops.get(name)


# ---------------------------------------------------------------- requests


class Requests(object):
    """Parses the generated requests/ package."""

    def __init__(self):
        self.by_op = {}

    def load(self):
        for fn in sorted(os.listdir(REQ_DIR)):
            if not fn.endswith(".go"):
                continue
            src = open(os.path.join(REQ_DIR, fn), encoding="utf-8").read()
            m = re.search(r"func (New\w+Request)\(([^)]*)\)", src)
            if not m:
                continue
            ctor = m.group(1)
            opname = ctor[len("New") : -len("Request")]
            ctor_params = []
            for part in m.group(2).split(","):
                part = part.strip()
                if not part or part.startswith("opts"):
                    continue
                bits = part.split()
                if len(bits) < 2:
                    continue
                ctor_params.append({"name": bits[0], "type": " ".join(bits[1:])})
            option_types = {}
            for om in re.finditer(
                r'cfg\.Params\["([^"]+)"\]\.\((\*?[A-Za-z0-9_.\[\]]+)\)', src
            ):
                option_types[om.group(1)] = om.group(2)
            self.by_op[opname] = {
                "ctor": ctor,
                "file": fn,
                "ctor_params": ctor_params,
                "option_types": option_types,
            }

    def get(self, opname):
        return self.by_op.get(opname)


# ---------------------------------------------------------------- models


class Models(object):
    """Parses field types out of the generated models/ package.

    Fields promoted from embedded structs (e.g. models.Picture embeds
    models.Shape) are flattened in so that `Picture.Left` resolves.
    """

    def __init__(self):
        self.fields = {}
        self.own = {}        # direct fields only
        self.embedded = {}   # struct -> [embedded struct names]
        self.interfaces = set()

    def load(self):
        for fn in sorted(os.listdir(MODEL_DIR)):
            if not fn.endswith(".go"):
                continue
            src = open(os.path.join(MODEL_DIR, fn), encoding="utf-8").read()
            self.interfaces.update(re.findall(r"type (\w+) interface", src))
            for sm in re.finditer(r"type (\w+) struct \{(.*?)\n\}", src, re.S):
                name, body = sm.group(1), sm.group(2)
                own, embeds = {}, []
                for line in body.splitlines():
                    line = line.strip()
                    if not line or line.startswith("//"):
                        continue
                    fm = re.match(r"(\w+)\s+([^\s`]+)\s+`", line) or re.match(
                        r"(\w+)\s+([^\s`]+)$", line
                    )
                    if fm:
                        own[fm.group(1)] = fm.group(2)
                    elif re.match(r"^\w+$", line):
                        embeds.append(line)  # embedded struct, promoted fields
                self.own[name] = own
                self.embedded[name] = embeds

        for name in self.own:
            self.fields[name] = self._flatten(name, set())

    def _flatten(self, name, seen):
        if name in seen:
            return {}
        seen = seen | {name}
        merged = dict(self.own.get(name, {}))
        for e in self.embedded.get(name, []):
            for k, v in self._flatten(e, seen).items():
                merged.setdefault(k, v)
        return merged

    def field_type(self, model, field):
        return self.fields.get(model, {}).get(field)


# ---------------------------------------------------------------- old tests


FUNC_RE = re.compile(r"^func (Test\w+)\(t \*testing\.T\) \{$", re.M)
DECL_RE = re.compile(
    r"^\s*(?:var\s+)?(\w+)\s*:?=\s*new\s*\(\s*(\w+)\s*\)\s*$"
)
ASSIGN_RE = re.compile(r"^\s*(\w+)\.([A-Za-z_]\w*)\s*=\s*(.+?)\s*$")
CALL_RE = re.compile(r"^.*CellsApi\.(\w+)\s*\(\s*(\w+)\s*\)\s*$")


def split_functions(src):
    """Yield (name, body_lines) for every top-level Test function."""
    lines = src.splitlines()
    i = 0
    while i < len(lines):
        m = FUNC_RE.match(lines[i])
        if not m:
            i += 1
            continue
        name = m.group(1)
        body = []
        depth = 1
        i += 1
        while i < len(lines) and depth > 0:
            line = lines[i]
            depth += line.count("{") - line.count("}")
            if depth > 0:
                body.append(line)
            i += 1
        yield name, body

    return


# ---------------------------------------------------------------- value conversion


class Converter(object):
    def __init__(self, spec, requests, models):
        self.spec = spec
        self.requests = requests
        self.models = models
        # local vars whose declared slice type must follow the model field type
        self.retarget = {}
        # longest first so `ColorFilterRequest` is not eaten by `Color`
        self.model_names = sorted(models.fields, key=len, reverse=True)
        alt = "|".join(re.escape(t) for t in self.model_names)
        self.re_new = re.compile(r"(?<![\w.])new\(\s*(%s)\s*\)" % alt)
        self.re_lit = re.compile(r"(?<![\w.])(%s)\s*\{" % alt)
        self.re_slice = re.compile(r"\[\](?<![\w.])(%s)\b(?!\s*\{)" % alt)
        self.stats = {
            "unmatched_field": [],
            "unresolved_op": [],
            "unresolved_model": [],
            "model_field_removed": [],
            "model_field_unsettable": [],
            "option_not_exposed": [],
            "test_not_portable": [],
        }

    # -- primitives ---------------------------------------------------

    @staticmethod
    def strip_cast(expr, casts=("int64", "int32", "int", "float64", "float32")):
        for c in casts:
            m = re.match(r"^%s\(\s*(.*?)\s*\)$" % c, expr)
            if m:
                return m.group(1)
        return expr

    def wrap_pointer(self, target, expr):
        """Wrap `expr` so it can be assigned to a field/option of `target` type."""
        if target in ("string", ""):
            # A string target never needs a pointer helper.
            return self.strip_cast(expr, ("int64", "int32", "int", "float64"))
        if target == "*bool":
            return "%s.BoolPtr(%s)" % (PKG, self.strip_cast(expr))
        if target == "*int32":
            return "%s.Int32Ptr(%s)" % (PKG, self.strip_cast(expr))
        if target == "*int64":
            return "%s.Int64Ptr(%s)" % (PKG, self.strip_cast(expr))
        if target == "*float64":
            return "%s.Float64Ptr(%s)" % (PKG, self.strip_cast(expr))
        if target == "*int":
            return "intPtr(%s)" % self.strip_cast(expr)
        if target in ("int", "int32", "int64", "float64", "float32"):
            # positional numeric argument: drop the legacy int64(...) cast
            return self.strip_cast(expr)
        if target == "[]interface{}":
            # e.g. CreatePivotTableRequest.PivotFieldColumns: the legacy setup
            # declares []int64, the v4.0 model wants []interface{}.
            m = re.match(r"^\[\](?:int64|int32|int|float64)\s*\{", expr)
            if m:
                return "[]interface{}{" + expr[m.end() :]
            if re.match(r"^\w+$", expr):
                self.retarget[expr] = "[]interface{}"
            return expr
        if target == "[]byte":
            # models.Color.A/R/G/B are declared []byte in the new model even
            # though the legacy ones were int64 -- see README "known differences".
            return "[]byte{%s}" % self.strip_cast(expr)
        # models / slices / anything else: pass through unchanged
        return expr

    def qualify(self, line):
        """Prefix unqualified model type references with the `models` package.

        Legacy setup lines such as `var options = []FontSetting{* optionsvalue0}`
        or `new(SaveOptions)` name the models unqualified; the new SDK splits
        them into the models package.
        """
        parts = re.split(r'("(?:[^"\\]|\\.)*")', line)  # keep string literals intact
        for i in range(0, len(parts), 2):
            s = self.re_new.sub(r"&models.\1{}", parts[i])
            s = self.re_lit.sub(r"models.\1{", s)
            s = self.re_slice.sub(r"[]models.\1", s)
            parts[i] = s
        return "".join(parts)

    def zero_value(self, target):
        if target.startswith("*") or target.startswith("[]"):
            return "nil"
        if target == "string":
            return '""'
        if target == "bool":
            return "false"
        if target in ("int", "int32", "int64", "float64", "float32"):
            return "0"
        return "nil"

    # -- one function -------------------------------------------------

    def convert_function(self, name, body, opname_hint):
        self.retarget = {}
        items = self.parse_body(body)
        opname = opname_hint
        for it in items:
            if it[0] == "call":
                opname = it[1]
                break
        op = self.spec.operation(opname)
        req = self.requests.get(opname) if opname else None
        if op is None or req is None:
            self.stats["unresolved_op"].append("%s -> %s" % (name, opname))
            return None, None

        # The request handed to the API call is the one and only "main" group;
        # other *Request types are plain body models.
        main_var = None
        for it in items:
            if it[0] == "call":
                main_var = it[2]

        if self.uses_multi_file(items, req):
            self.stats["test_not_portable"].append(name)
            return None, (MULTI_FILE_SKIP, MULTI_FILE_REASON)

        out = []
        for it in items:
            if it[0] == "line":
                out.append(it[1])
            elif it[0] == "group":
                _, var, mtype, assigns = it
                if var == main_var:
                    out.extend(self.emit_main(assigns, req, op))
                elif mtype == "UploadFileRequest":
                    out.extend(self.emit_upload(assigns))
                else:
                    out.extend(self.emit_model(var, mtype, assigns))
            else:  # call
                out.append("")
                out.extend(self.emit_call(name, opname, main_var))
        return self.apply_retarget(out), None

    def apply_retarget(self, lines):
        """Rewrite the declared slice type of locals whose model field target
        turned out to be a different (interface) slice type."""
        for var, newtype in self.retarget.items():
            pat = r"(\b%s\s*=\s*)(?:\[\](?:int64|int32|int|float64))\s*\{" % re.escape(var)
            for i, line in enumerate(lines):
                lines[i] = re.sub(pat, r"\1%s{" % newtype, line)
        return lines

    def uses_multi_file(self, items, req):
        """True when a request field is fed a map of local files.

        PostAssemble is the only such operation: the legacy request took
        `File map[string]string` and uploaded every entry as its own multipart
        part, while the v4.0 request exposes a single `File string`.
        """
        maps = set()
        for it in items:
            if it[0] == "line":
                m = re.search(r"(\w+)\s+map\[string\]string", it[1]) or re.search(
                    r"(\w+)\s*:=\s*make\(map\[string\]string\)", it[1]
                )
                if m:
                    maps.add(m.group(1))
        if not maps:
            return False
        for it in items:
            if it[0] != "group":
                continue
            for _field, expr in it[3]:
                if expr.strip() in maps and "string" in [
                    p["type"] for p in req["ctor_params"]
                ]:
                    return True
        return False

    def parse_body(self, body):
        """Split a function body into an ordered item list, so that the emitted
        code keeps the original declaration / assignment order.

        Items are ("line", text) for passthrough setup, ("group", var, type,
        assignments) for declarations, and ("call", opName, var) for the API
        invocation (whose result-checking block is regenerated).
        """
        items = []
        groups = {}  # var -> item, so later field assignments attach to it
        call_op = None

        for line in body:
            call = CALL_RE.match(line)
            if call:
                op, var = call.group(1), call.group(2)
                # `GetBaseTest().CellsApi.UploadFile(x)` written as a bare
                # statement is the setup upload -- it is already folded into the
                # group's mustUploadFile. The same operation written as an
                # assignment (`_, httpResponse, err := ...`) is the real subject
                # of TestFileController_UploadFile and must be kept.
                if "=" not in line:
                    continue
                call_op = op
                items.append(("call", op, var))
                continue
            decl = DECL_RE.match(line)
            if decl:
                item = ("group", decl.group(1), decl.group(2), [])
                groups[decl.group(1)] = item
                items.append(item)
                continue
            assign = ASSIGN_RE.match(line)
            if assign and assign.group(1) in groups:
                groups[assign.group(1)][3].append((assign.group(2), assign.group(3)))
                continue
            if self.is_check_block_line(line):
                continue
            # keep meaningful setup lines
            if line.strip():
                items.append(("line", self.qualify(line).rstrip()))

        # drop trailing blanks
        while items and items[-1][0] == "line" and not items[-1][1].strip():
            items.pop()
        return items

    @staticmethod
    def is_check_block_line(line):
        s = line.strip()
        return (
            s.startswith("if err != nil {")
            or s.startswith("} else if httpResponse")
            or s.startswith("} else {")
            or s == "}"
            or s.startswith("t.Error(")
            or s.startswith("t.Fail()")
            or s.startswith("fmt.Printf(")
            or s.startswith("_, httpResponse")
            or s.startswith("_, err")
        )

    def emit_upload(self, assigns):
        fields = {f: v for f, v in assigns}
        path = fields.get("Path", '""')
        files = fields.get("UploadFiles", '""')
        storage = fields.get("StorageName", '""')
        return [
            "if err := mustUploadFile(t, %s, %s, %s); err != nil {" % (path, files, storage),
            "\tt.Fatal(err)",
            "}",
            "",
        ]

    def emit_model(self, var, mtype, assigns):
        if not self.models.fields.get(mtype):
            self.stats["unresolved_model"].append(mtype)
        out = ["%s := &models.%s{}" % (var, mtype)]
        for field, expr in assigns:
            ftype = self.models.field_type(mtype, field)
            if ftype is None and field.endswith("_"):
                ftype = self.models.field_type(mtype, field[:-1])
                field = field[:-1]
            if ftype is None:
                # Field dropped from the new model -- skip it and report.
                self.stats["model_field_removed"].append("%s.%s" % (mtype, field))
                continue
            if ftype.startswith("*") and ftype[1:] in self.models.interfaces:
                # Pointer to an interface: nothing can ever satisfy it, so the
                # legacy value has no place to go. See README "known differences".
                self.stats["model_field_unsettable"].append(
                    "%s.%s (%s)" % (mtype, field, ftype)
                )
                continue
            out.append("%s.%s = %s" % (var, field, self.wrap_pointer(ftype, expr)))
        return out

    def emit_main(self, assigns, req, op):
        spec_params = op["params"]
        # map normalized name -> spec param
        by_norm = {norm(p["name"]): p for p in spec_params}
        # values gathered from the old test
        values = {}
        leftovers = []
        def lookup(field):
            key = norm(field)
            p = by_norm.get(key)
            if p is None and key in ALIASES:
                p = by_norm.get(norm(ALIASES[key]))
            return p

        for field, expr in assigns:
            p = lookup(field)
            if p is None:
                self.stats["unmatched_field"].append("%s (op param)" % field)
                continue
            values[p["name"]] = expr

        # positional args, in the constructor's own order
        args = []
        used = set()
        for cp in req["ctor_params"]:
            p = by_norm.get(norm(cp["name"]))
            if p is not None:
                used.add(p["name"])
            expr = values.get(p["name"]) if p else None
            if expr is None:
                args.append(self.zero_value(cp["type"]))
            else:
                args.append(self.wrap_pointer(cp["type"], expr))

        # remaining values become options
        opts = []
        for field, expr in assigns:
            p = lookup(field)
            if p is None or p["name"] in used:
                continue
            otype = req["option_types"].get(p["name"])
            if otype is None:
                # The operation accepts the parameter but the generated request
                # never reads it -- the value is silently un-settable.
                self.stats["option_not_exposed"].append(
                    "%s.%s (op %s)" % (field, p["name"], op["name"])
                )
                continue
            opts.append(
                'requests.WithCommonParameter("%s", %s)'
                % (p["name"], self.wrap_pointer(otype, expr))
            )

        lines = ["request := requests.%s(" % req["ctor"]]
        for arg in args + opts:
            lines.append("\t%s," % arg)
        lines.append(")")
        return lines

    def emit_call(self, name, opname, var):
        return [
            "_, err := %s.DoChecked(ctx, GetBaseTest().Client, %s)" % (PKG, var),
            "if err != nil {",
            "\tt.Error(err)",
            "} else {",
            '\tfmt.Printf("%%d\\t%s \\n", GetBaseTest().GetTestNumber())' % name,
            "}",
        ]


# ---------------------------------------------------------------- file emit


HEADER = """package newtests

import (
\t"context"
\t"fmt"
\t"testing"

\t"asposecellscloud"
%s\t"asposecellscloud/requests"
)
"""


def emit_file(functions):
    body = []
    uses_models = any("models." in ln for fn in functions for ln in fn)
    imports = '\t"asposecellscloud/models"\n' if uses_models else ""
    for fn in functions:
        body.append("")
        body.extend(fn)
    return HEADER % imports + "\n".join(body) + "\n"


MULTI_FILE_SKIP = "multi-file upload is not expressible in the v4.0 request model"

MULTI_FILE_REASON = (
    "PostAssemble and PostMerge upload several workbooks at once: the legacy "
    "request took File map[string]string and sent every entry as its own "
    "multipart part. The v4.0 request exposes a single File string, so the "
    "multi-file assembly cannot be expressed."
)


def main():
    spec = Spec()
    requests = Requests()
    requests.load()
    models = Models()
    models.load()
    conv = Converter(spec, requests, models)

    if not os.path.isdir(OUT_DIR):
        os.makedirs(OUT_DIR)

    total = 0
    for fn in sorted(os.listdir(SRC_DIR)):
        if not fn.endswith("_test.go") or fn == "base_test.go":
            continue
        src = open(os.path.join(SRC_DIR, fn), encoding="utf-8").read()
        out_funcs = []
        for name, body in split_functions(src):
            sig = "func %s(t *testing.T) {" % name
            lines, stub = conv.convert_function(name, body, None)
            if lines is None and stub is None:
                print("SKIP %s" % name, file=sys.stderr)
                continue
            if stub is not None:
                short, long = stub
                lines = ["// " + l for l in textwrap.wrap(long, 76)]
                lines += [sig, "\tt.Skip(%s)" % json.dumps(short), "}"]
            else:
                lines = [sig, "\tctx := context.Background()"] + [
                    "\t" + l if l.strip() else l for l in lines
                ] + ["}"]
            out_funcs.append(lines)
            total += 1
        with open(os.path.join(OUT_DIR, fn), "w", encoding="utf-8") as fh:
            fh.write(emit_file(out_funcs))

    # the emitter writes plain Go; let gofmt own the layout
    if subprocess.call(["gofmt", "-w", OUT_DIR]) != 0:
        print("warning: gofmt failed", file=sys.stderr)

    print("converted functions: %d" % total)
    for key, items in conv.stats.items():
        distinct = sorted(set(items))
        print("%s: %d (%d distinct)" % (key, len(items), len(distinct)))
        for it in distinct[:40]:
            print("   ", it)


if __name__ == "__main__":
    main()
