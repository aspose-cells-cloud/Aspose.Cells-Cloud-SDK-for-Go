//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func main() {
	ctx := context.Background()

	// If no environment variables are configured, obtain the ClientId and
	// ClientSecret from https://dashboard.aspose.cloud/#/applications and pass
	// them to NewAsposeCellsCloudClient explicitly:
	//   asposecellscloud.NewAsposeCellsCloudClient("YourClientId", "YourClientSecret", "")
	client := asposecellscloud.NewAsposeCellsCloudClient(
		os.Getenv("ProductClientId"),
		os.Getenv("ProductClientSecret"),
		os.Getenv("CellsCloudApiBaseUrl"),
	)

	const input = "BookText.xlsx"

	// Trim leading/trailing characters and repeated spaces in the whole file.
	run(ctx, client, "out.xlsx",
		requests.NewTrimCharacterRequest(input,
			requests.WithCommonParameter("trimLeading", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("trimTrailing", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("trimSpaceBetweenWordTo1", asposecellscloud.BoolPtr(true)),
		))

	// Trim a single cell range.
	run(ctx, client, "out1.xlsx",
		requests.NewTrimCharacterRequest(input,
			requests.WithCommonParameter("trimLeading", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("trimTrailing", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("trimSpaceBetweenWordTo1", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "J6:J6"),
		))

	// Trim trailing characters and extra line breaks in a single cell.
	run(ctx, client, "out2.xlsx",
		requests.NewTrimCharacterRequest(input,
			requests.WithCommonParameter("trimLeading", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("trimTrailing", asposecellscloud.BoolPtr(false)),
			requests.WithCommonParameter("trimSpaceBetweenWordTo1", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("removeExtraLineBreaks", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "J7:J7"),
		))

	// Change the word case of a cell, chaining each result into the next call.
	run(ctx, client, "out3.xlsx",
		requests.NewUpdateWordCaseRequest(input, "UpperCase",
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "J7:J7"),
		))
	run(ctx, client, "out4.xlsx",
		requests.NewUpdateWordCaseRequest("out3.xlsx", "LowerCase",
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "J7:J7"),
		))
	run(ctx, client, "out5.xlsx",
		requests.NewUpdateWordCaseRequest("out4.xlsx", "ProperCase",
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "J7:J7"),
		))
	run(ctx, client, "out6.xlsx",
		requests.NewUpdateWordCaseRequest("out5.xlsx", "SentenceCase",
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "J7:J7"),
		))

	// Add text at the beginning of the workbook.
	run(ctx, client, "out7.xlsx",
		requests.NewAddTextRequest("AtTheBeginning", input, "Hello, "))

	// Add text after the selected text.
	run(ctx, client, "out8.xlsx",
		requests.NewAddTextRequest("AfterText", input, "!",
			requests.WithCommonParameter("selectText", "hello world"),
		))

	// Add text at the end of a cell range.
	run(ctx, client, "out9.xlsx",
		requests.NewAddTextRequest("AtTheEnd", input, "!",
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "A1:J7"),
		))

	// Convert numbers to text.
	run(ctx, client, "out10.xlsx",
		requests.NewConvertTextRequest("ConvertNumberToText", input))

	// Convert characters.
	run(ctx, client, "out11.xlsx",
		requests.NewConvertTextRequest("ConvertCharacters", input,
			requests.WithCommonParameter("sourceCharacters", "1"),
			requests.WithCommonParameter("targetCharacters", "Hello World"),
			requests.WithCommonParameter("worksheet", "Text"),
			requests.WithCommonParameter("range", "A1:J7"),
		))

	// Convert line breaks.
	run(ctx, client, "out12.xlsx",
		requests.NewConvertTextRequest("ConvertLinebreak", input,
			requests.WithCommonParameter("sourceCharacters", ""),
			requests.WithCommonParameter("targetCharacters", "Linebreak"),
		))

	// Convert whitespace.
	run(ctx, client, "out13.xlsx",
		requests.NewConvertTextRequest("ConvertWriteSpace", input,
			requests.WithCommonParameter("sourceCharacters", ""),
			requests.WithCommonParameter("targetCharacters", "WriteSpace"),
		))
}

// run executes a local content request and writes the returned workbook to
// outPath.
func run(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, outPath string, req asposecellscloud.RequestOption) {
	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(outPath, resp.Body, 0644); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
