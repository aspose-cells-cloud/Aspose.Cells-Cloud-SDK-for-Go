package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type SwapRangeRequest struct {
	range1          string
	range2          string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	worksheet1      string
	worksheet2      string

	outPath        string
	outStorageName string
	password       string
	region         string

	extraQueryParameters map[string]string
}

func NewSwapRangeRequest(range1 string, range2 string, Spreadsheet string, worksheet1 string, worksheet2 string, opts ...Option) *SwapRangeRequest {
	req := &SwapRangeRequest{
		range1:      range1,
		range2:      range2,
		Spreadsheet: Spreadsheet,
		worksheet1:  worksheet1,
		worksheet2:  worksheet2,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
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

func (request *SwapRangeRequest) Validate() error {
	if request.range1 == "" {
		return fmt.Errorf("required request parameter %q is missing", "range1")
	}

	if request.range2 == "" {
		return fmt.Errorf("required request parameter %q is missing", "range2")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.worksheet1 == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet1")
	}

	if request.worksheet2 == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet2")
	}

	return nil
}

func (request *SwapRangeRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *SwapRangeRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SwapRangeRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SwapRangeRequest) GetMethod() string {
	return "PUT"
}

func (request *SwapRangeRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *SwapRangeRequest) GetPath() string {
	localVarPath := "/v4.0/cells/swap/range"
	return localVarPath
}

func (request *SwapRangeRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet1", fmt.Sprintf("%v", request.worksheet1))
	localVarQueryParams.Add("range1", fmt.Sprintf("%v", request.range1))
	localVarQueryParams.Add("worksheet2", fmt.Sprintf("%v", request.worksheet2))
	localVarQueryParams.Add("range2", fmt.Sprintf("%v", request.range2))
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
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

func (request *SwapRangeRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SwapRangeRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
