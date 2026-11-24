package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type ConvertWorksheetToHtmlTableRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	worksheet       string

	password string
	region   string

	extraQueryParameters map[string]string
}

func NewConvertWorksheetToHtmlTableRequest(Spreadsheet string, worksheet string, opts ...Option) *ConvertWorksheetToHtmlTableRequest {
	req := &ConvertWorksheetToHtmlTableRequest{
		Spreadsheet: Spreadsheet,
		worksheet:   worksheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if len(cfg.extraQueryParams) > 0 {
		if req.extraQueryParameters == nil {
			req.extraQueryParameters = make(map[string]string)
		}
		for k, v := range cfg.extraQueryParams {
			req.extraQueryParameters[k] = v
		}
	}

	return req
}

func (request *ConvertWorksheetToHtmlTableRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *ConvertWorksheetToHtmlTableRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *ConvertWorksheetToHtmlTableRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ConvertWorksheetToHtmlTableRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ConvertWorksheetToHtmlTableRequest) GetMethod() string {
	return "PUT"
}

func (request *ConvertWorksheetToHtmlTableRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *ConvertWorksheetToHtmlTableRequest) GetPath() string {
	localVarPath := "/v4.0/cells/convert/worksheet/html-table"
	return localVarPath
}

func (request *ConvertWorksheetToHtmlTableRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *ConvertWorksheetToHtmlTableRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ConvertWorksheetToHtmlTableRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
