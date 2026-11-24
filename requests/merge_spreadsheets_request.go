package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type MergeSpreadsheetsRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	fontsLocation   string
	mergeInOneSheet *bool
	outFormat       string
	outPath         string
	outStorageName  string
	password        string
	region          string

	extraQueryParameters map[string]string
}

func NewMergeSpreadsheetsRequest(Spreadsheet string, opts ...Option) *MergeSpreadsheetsRequest {
	req := &MergeSpreadsheetsRequest{
		Spreadsheet: Spreadsheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["fontsLocation"].(string); ok {
		req.fontsLocation = val
	}
	if val, ok := cfg.Params["mergeInOneSheet"].(*bool); ok {
		req.mergeInOneSheet = val
	}
	if val, ok := cfg.Params["outFormat"].(string); ok {
		req.outFormat = val
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

func (request *MergeSpreadsheetsRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *MergeSpreadsheetsRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *MergeSpreadsheetsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *MergeSpreadsheetsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *MergeSpreadsheetsRequest) GetMethod() string {
	return "PUT"
}

func (request *MergeSpreadsheetsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *MergeSpreadsheetsRequest) GetPath() string {
	localVarPath := "/v4.0/cells/merge/spreadsheet"
	return localVarPath
}

func (request *MergeSpreadsheetsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.outFormat != "" {
		localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
	}
	if request.mergeInOneSheet != nil {
		localVarQueryParams.Add("mergeInOneSheet", fmt.Sprintf("%v", *request.mergeInOneSheet))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.fontsLocation != "" {
		localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
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

func (request *MergeSpreadsheetsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *MergeSpreadsheetsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
