# newtests 集成测试报告

> 本报告由 `newtests/run-tests.ps1` 自动生成，数据来源为本次运行的原始 `go test -v` 日志与 `test-logs/*.tsv` 记录。

## 一、运行环境

| 项目 | 值 |
|---|---|
| 测试套件 | newtests（v4.0 新模型 API 集成测试） |
| API 基地址 | `https://api.aspose.cloud` |
| SDK 版本 | `v26.8.0` |
| 单文件超时 | `20m` |
| 开始时间 | 2026-09-22 23:06:05 |
| 结束时间 | 2026-09-22 23:47:17 |
| 总耗时 | 2471.8s |
| 日志目录 | `test-logs` |
| 失败用例重跑 | 已启用 |

## 二、总览

| 指标 | 数值 |
|---|---|
| 测试文件 | 46 |
| 用例总数（含跳过） | 509 |
| 通过 | 489 |
| 失败 | 7 |
| 跳过 | 13 |
| 编译失败文件 | 0 |
| 超时文件 | 0 |
| 通过率 | 98.59% |

**结论：本次运行存在失败项**，明细见第三、四节。

## 三、逐文件结果

| 文件 | 用例 | 通过 | 失败 | 跳过 | 状态 | 耗时 | 调用的 API 版本 |
|---|---:|---:|---:|---:|---|---:|---|
| `api_cells_addtext_test.go` | 4 | 4 | 0 | 0 | PASS | 14.7s | v4.0 |
| `api_cells_autofiltercontroller_test.go` | 13 | 13 | 0 | 0 | PASS | 23.4s | v3.0 |
| `api_cells_batchcontroller_test.go` | 5 | 5 | 0 | 0 | PASS | 25.7s | v3.0 |
| `api_cells_calculate_test.go` | 2 | 2 | 0 | 0 | PASS | 5.3s | v4.0 |
| `api_cells_cellscontroller_test.go` | 40 | 40 | 0 | 0 | PASS | 78.2s | v3.0 |
| `api_cells_cellsstatuscontroller_test.go` | 2 | 2 | 0 | 0 | PASS | 16.6s | v3.0 |
| `api_cells_chartareacontroller_test.go` | 3 | 3 | 0 | 0 | PASS | 9.5s | v3.0 |
| `api_cells_chartscontroller_test.go` | 14 | 14 | 0 | 0 | PASS | 31.4s | v3.0 |
| `api_cells_conditionalformattingscontroller_test.go` | 9 | 9 | 0 | 0 | PASS | 32.3s | v3.0 |
| `api_cells_conversion_test.go` | 31 | 31 | 0 | 0 | PASS | 140.8s | v4.0 |
| `api_cells_conversion30_test.go` | 70 | 68 | 2 | 0 | FAIL | 518.3s | v3.0 |
| `api_cells_converttext_test.go` | 3 | 3 | 0 | 0 | PASS | 12.9s | v4.0 |
| `api_cells_dataprocessingcontroller_test.go` | 4 | 4 | 0 | 0 | PASS | 10.4s | v3.0 |
| `api_cells_extracttext_test.go` | 3 | 3 | 0 | 0 | PASS | 7.1s | v4.0 |
| `api_cells_filecontroller_test.go` | 3 | 3 | 0 | 0 | PASS | 22.4s | v4.0 |
| `api_cells_foldercontroller_test.go` | 5 | 5 | 0 | 0 | PASS | 17.8s | v4.0 |
| `api_cells_hypelinkscontroller_test.go` | 6 | 6 | 0 | 0 | PASS | 16.5s | v3.0 |
| `api_cells_importdata_test.go` | 1 | 1 | 0 | 0 | PASS | 21.6s | v4.0 |
| `api_cells_lightcells_test.go` | 72 | 55 | 4 | 13 | FAIL | 693.3s | v3.0 |
| `api_cells_listobjectscontroller_test.go` | 13 | 13 | 0 | 0 | PASS | 27.2s | v3.0 |
| `api_cells_management_test.go` | 7 | 7 | 0 | 0 | PASS | 33.0s | v4.0 |
| `api_cells_merger_test.go` | 1 | 1 | 0 | 0 | PASS | 18.1s | v4.0 |
| `api_cells_oleobjectscontroller_test.go` | 6 | 6 | 0 | 0 | PASS | 15.2s | v3.0 |
| `api_cells_pagebreakscontroller_test.go` | 10 | 10 | 0 | 0 | PASS | 19.6s | v3.0 |
| `api_cells_pagesetupcontroller_test.go` | 9 | 9 | 0 | 0 | PASS | 20.3s | v3.0 |
| `api_cells_picturescontroller_test.go` | 6 | 6 | 0 | 0 | PASS | 16.6s | v3.0 |
| `api_cells_pivottablescontroller_test.go` | 19 | 19 | 0 | 0 | PASS | 26.8s | v3.0 |
| `api_cells_propertiescontroller_test.go` | 5 | 5 | 0 | 0 | PASS | 12.0s | v3.0 |
| `api_cells_protection_test.go` | 2 | 2 | 0 | 0 | PASS | 6.4s | v4.0 |
| `api_cells_rangescontroller_test.go` | 13 | 13 | 0 | 0 | PASS | 26.5s | v3.0 |
| `api_cells_removecharacters_test.go` | 5 | 5 | 0 | 0 | PASS | 14.0s | v4.0 |
| `api_cells_replacer_test.go` | 3 | 3 | 0 | 0 | PASS | 9.5s | v4.0 |
| `api_cells_searcher_test.go` | 8 | 8 | 0 | 0 | PASS | 14.5s | v4.0 |
| `api_cells_shapescontroller_test.go` | 8 | 8 | 0 | 0 | PASS | 16.7s | v3.0 |
| `api_cells_sparklinegroupscontroller_test.go` | 6 | 6 | 0 | 0 | PASS | 10.7s | v3.0 |
| `api_cells_splitter_test.go` | 3 | 2 | 1 | 0 | FAIL | 51.0s | v4.0 |
| `api_cells_splittext_test.go` | 1 | 1 | 0 | 0 | PASS | 4.9s | v4.0 |
| `api_cells_storagecontroller_test.go` | 4 | 4 | 0 | 0 | PASS | 12.7s | v4.0 |
| `api_cells_textprocessingcontroller_test.go` | 3 | 3 | 0 | 0 | PASS | 8.6s | v3.0 |
| `api_cells_transform_test.go` | 3 | 3 | 0 | 0 | PASS | 9.4s | v4.0 |
| `api_cells_trimspreadsheet_test.go` | 2 | 2 | 0 | 0 | PASS | 6.2s | v4.0 |
| `api_cells_updatewordcase_test.go` | 3 | 3 | 0 | 0 | PASS | 13.7s | v4.0 |
| `api_cells_workbookcontroller_test.go` | 32 | 32 | 0 | 0 | PASS | 66.5s | v3.0 |
| `api_cells_worksheetcontroller_test.go` | 39 | 39 | 0 | 0 | PASS | 65.5s | v3.0 |
| `api_cells_worksheetvalidationscontroller_test.go` | 6 | 6 | 0 | 0 | PASS | 14.3s | v3.0 |
| `api_cells_xmlcontroller_test.go` | 2 | 2 | 0 | 0 | PASS | 8.6s | v3.0 |

## 四、失败用例明细

| # | 文件 | 用例 | 单独重跑 | 原因分类 | 调用的 API 版本 | 错误摘要 |
|---:|---|---|---|---|---|---|
| 1 | `api_cells_conversion30_test.go` | `TestConversion30_GetWorkbookFormat_tif` | PASS(重跑通过) | 偶发失败(重跑通过) | v3.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |
| 2 | `api_cells_conversion30_test.go` | `TestConversion30_ConvertWorkbook_tif` | FAIL | 网络错误 | v3.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |
| 3 | `api_cells_lightcells_test.go` | `TestLightCells_PostSplit_html` | FAIL | 网络错误 | v3.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |
| 4 | `api_cells_lightcells_test.go` | `TestLightCells_PostSplit_xps` | FAIL | 网络错误 | v3.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |
| 5 | `api_cells_lightcells_test.go` | `TestLightCells_PostSplit_pptx` | FAIL | 网络错误 | v3.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |
| 6 | `api_cells_lightcells_test.go` | `TestLightCells_PostExport_html_worksheet` | PASS(重跑通过) | 偶发失败(重跑通过) | v3.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |
| 7 | `api_cells_splitter_test.go` | `TestSplitter_SplitLocalFile` | PASS(重跑通过) | 偶发失败(重跑通过) | v4.0 | sdk error [code=-1]: request execution failed: context deadline exceeded (Client.Timeout or context cancellation while reading body) |

## 五、失败原因分类统计

| 原因分类 | 数量 |
|---|---:|
| 网络错误 | 4 |
| 偶发失败(重跑通过) | 3 |

## 六、API 版本维度分析

每个生成的 request 现在携带其 operation 自己的 API 版本前缀（v3.0 或 v4.0，取自规范中该 operation 的 `APIVersion`）。
下表按用例调用的 API 版本统计通过率，用于判断失败是否集中在某一个版本上。

| 用例调用的 API 版本 | 用例数 | 通过 | 失败 | 通过率 |
|---|---:|---:|---:|---:|
| v3.0 | 402 | 396 | 6 | 98.51% |
| v4.0 | 94 | 93 | 1 | 98.94% |

## 七、分析结论与建议

- 共执行 46 个文件 / 509 个用例，通过 489，失败 7，跳过 13，通过率 **98.59%**。
- 失败最集中的原因分类是 **网络错误**（4 个用例），应优先排查。
- 仅调用 v3.0 接口的用例 402 个，失败率 1.49%；仅调用 v4.0 接口的用例 94 个，失败率 1.06%。
- 失败中有 7 个属于**传输层问题**（4 个网络错误 + 3 个重跑即通过），表现为 `context deadline exceeded` / `TLS handshake timeout`，属于服务端或网络抖动，而非接口契约或请求模型缺陷；建议先按文件重跑以剔除这些噪声，再分析剩余失败。
- 有 3 个用例在单独重跑时通过，说明存在**偶发性（服务器抖动）**或共享远程文件状态的因素，建议确认测试之间的远程目录是否需要隔离。
- 有 13 个用例被跳过（`t.Skip`），多为多文件上传受限的 LightCells 组装/合并用例，不计入通过率。
- 详细日志：`test-logs/results.tsv`（逐文件）、`test-logs/failures.tsv`（逐失败用例）、`test-logs/<文件>.log`（原始输出）。

## 八、跳过的用例

| 文件 | 用例 |
|---|---|
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_csv` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_html` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_ods` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_pdf` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_md` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_svg` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_docx` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_pptx` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostAssemble_json` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostMerge_html_true` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostMerge_pdf_true` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostMerge_xlsx_true` |
| `api_cells_lightcells_test.go` | `TestLightCells_PostMerge_json_false` |

---

_报告由 `newtests/run-tests.ps1` 生成。_
