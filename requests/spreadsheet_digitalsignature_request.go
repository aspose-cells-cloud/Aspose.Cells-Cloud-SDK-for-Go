package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type SpreadsheetDigitalsignatureRequest struct {
	password        string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	outPath        string
	outStorageName string
	region         string

	extraQueryParameters map[string]string
}

func NewSpreadsheetDigitalsignatureRequest(password string, Spreadsheet string, opts ...Option) *SpreadsheetDigitalsignatureRequest {
	req := &SpreadsheetDigitalsignatureRequest{
		password:    password,
		Spreadsheet: Spreadsheet,
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

func (request *SpreadsheetDigitalsignatureRequest) Validate() error {
	if request.password == "" {
		return fmt.Errorf("required request parameter %q is missing", "password")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *SpreadsheetDigitalsignatureRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *SpreadsheetDigitalsignatureRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SpreadsheetDigitalsignatureRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SpreadsheetDigitalsignatureRequest) GetMethod() string {
	return "PUT"
}

func (request *SpreadsheetDigitalsignatureRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *SpreadsheetDigitalsignatureRequest) GetPath() string {
	localVarPath := "/v4.0/cells/digitalsignature/spreadsheet"
	return localVarPath
}

func (request *SpreadsheetDigitalsignatureRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *SpreadsheetDigitalsignatureRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SpreadsheetDigitalsignatureRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
