package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type ConvertWorksheetToJsonRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	worksheet       string

	AutoColumnsFit *bool
	AutoRowsFit    *bool
	fontsLocation  string
	outPath        string
	outStorageName string
	password       string
	region         string

	extraQueryParameters map[string]string
}

func NewConvertWorksheetToJsonRequest(Spreadsheet string, worksheet string, opts ...Option) *ConvertWorksheetToJsonRequest {
	req := &ConvertWorksheetToJsonRequest{
		Spreadsheet: Spreadsheet,
		worksheet:   worksheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["AutoColumnsFit"].(*bool); ok {
		req.AutoColumnsFit = val
	}
	if val, ok := cfg.Params["AutoRowsFit"].(*bool); ok {
		req.AutoRowsFit = val
	}
	if val, ok := cfg.Params["fontsLocation"].(string); ok {
		req.fontsLocation = val
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

func (request *ConvertWorksheetToJsonRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *ConvertWorksheetToJsonRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *ConvertWorksheetToJsonRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ConvertWorksheetToJsonRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ConvertWorksheetToJsonRequest) GetMethod() string {
	return "PUT"
}

func (request *ConvertWorksheetToJsonRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *ConvertWorksheetToJsonRequest) GetPath() string {
	localVarPath := "/v4.0/cells/convert/worksheet/json"
	return localVarPath
}

func (request *ConvertWorksheetToJsonRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.fontsLocation != "" {
		localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
	}
	if request.AutoRowsFit != nil {
		localVarQueryParams.Add("AutoRowsFit", fmt.Sprintf("%v", *request.AutoRowsFit))
	}
	if request.AutoColumnsFit != nil {
		localVarQueryParams.Add("AutoColumnsFit", fmt.Sprintf("%v", *request.AutoColumnsFit))
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

func (request *ConvertWorksheetToJsonRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ConvertWorksheetToJsonRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
