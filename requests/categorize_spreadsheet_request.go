package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type CategorizeSpreadsheetRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	targetColumn    string

	newColumnName string
	password      string
	region        string
	sheetName     string

	extraQueryParameters map[string]string
}

func NewCategorizeSpreadsheetRequest(Spreadsheet string, targetColumn string, opts ...Option) *CategorizeSpreadsheetRequest {
	req := &CategorizeSpreadsheetRequest{
		Spreadsheet:  Spreadsheet,
		targetColumn: targetColumn,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["newColumnName"].(string); ok {
		req.newColumnName = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["sheetName"].(string); ok {
		req.sheetName = val
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

func (request *CategorizeSpreadsheetRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.targetColumn == "" {
		return fmt.Errorf("required request parameter %q is missing", "targetColumn")
	}

	return nil
}

func (request *CategorizeSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *CategorizeSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CategorizeSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CategorizeSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *CategorizeSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *CategorizeSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/ai/categorize/spreadsheet"
	return localVarPath
}

func (request *CategorizeSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("targetColumn", fmt.Sprintf("%v", request.targetColumn))
	if request.sheetName != "" {
		localVarQueryParams.Add("sheetName", fmt.Sprintf("%v", request.sheetName))
	}
	if request.newColumnName != "" {
		localVarQueryParams.Add("newColumnName", fmt.Sprintf("%v", request.newColumnName))
	}
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

func (request *CategorizeSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *CategorizeSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
