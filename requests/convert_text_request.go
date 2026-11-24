package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type ConvertTextRequest struct {
	convertTextType string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	outPath          string
	outStorageName   string
	password         string
	_range           string
	region           string
	sourceCharacters string
	targetCharacters string
	worksheet        string

	extraQueryParameters map[string]string
}

func NewConvertTextRequest(convertTextType string, Spreadsheet string, opts ...Option) *ConvertTextRequest {
	req := &ConvertTextRequest{
		convertTextType: convertTextType,
		Spreadsheet:     Spreadsheet,
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
	if val, ok := cfg.Params["range"].(string); ok {
		req._range = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["sourceCharacters"].(string); ok {
		req.sourceCharacters = val
	}
	if val, ok := cfg.Params["targetCharacters"].(string); ok {
		req.targetCharacters = val
	}
	if val, ok := cfg.Params["worksheet"].(string); ok {
		req.worksheet = val
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

func (request *ConvertTextRequest) Validate() error {
	if request.convertTextType == "" {
		return fmt.Errorf("required request parameter %q is missing", "convertTextType")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *ConvertTextRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *ConvertTextRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ConvertTextRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ConvertTextRequest) GetMethod() string {
	return "PUT"
}

func (request *ConvertTextRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *ConvertTextRequest) GetPath() string {
	localVarPath := "/v4.0/cells/content/convert/text"
	return localVarPath
}

func (request *ConvertTextRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("convertTextType", fmt.Sprintf("%v", request.convertTextType))
	if request.sourceCharacters != "" {
		localVarQueryParams.Add("sourceCharacters", fmt.Sprintf("%v", request.sourceCharacters))
	}
	if request.targetCharacters != "" {
		localVarQueryParams.Add("targetCharacters", fmt.Sprintf("%v", request.targetCharacters))
	}
	if request.worksheet != "" {
		localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	}
	if request._range != "" {
		localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
	}
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

func (request *ConvertTextRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ConvertTextRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
