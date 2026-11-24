package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type ExtractTextRequest struct {
	extractTextType  string
	outPositionRange string
	Spreadsheet      string
	SpreadsheetData  []byte
	SpreadsheetName  string

	afterPosition  *int
	afterText      string
	beforePosition *int
	beforeText     string
	outPath        string
	outStorageName string
	password       string
	_range         string
	region         string
	worksheet      string

	extraQueryParameters map[string]string
}

func NewExtractTextRequest(extractTextType string, outPositionRange string, Spreadsheet string, opts ...Option) *ExtractTextRequest {
	req := &ExtractTextRequest{
		extractTextType:  extractTextType,
		outPositionRange: outPositionRange,
		Spreadsheet:      Spreadsheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["afterPosition"].(*int); ok {
		req.afterPosition = val
	}
	if val, ok := cfg.Params["afterText"].(string); ok {
		req.afterText = val
	}
	if val, ok := cfg.Params["beforePosition"].(*int); ok {
		req.beforePosition = val
	}
	if val, ok := cfg.Params["beforeText"].(string); ok {
		req.beforeText = val
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

func (request *ExtractTextRequest) Validate() error {
	if request.extractTextType == "" {
		return fmt.Errorf("required request parameter %q is missing", "extractTextType")
	}

	if request.outPositionRange == "" {
		return fmt.Errorf("required request parameter %q is missing", "outPositionRange")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *ExtractTextRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *ExtractTextRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ExtractTextRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ExtractTextRequest) GetMethod() string {
	return "PUT"
}

func (request *ExtractTextRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *ExtractTextRequest) GetPath() string {
	localVarPath := "/v4.0/cells/content/extract/text"
	return localVarPath
}

func (request *ExtractTextRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("extractTextType", fmt.Sprintf("%v", request.extractTextType))
	localVarQueryParams.Add("outPositionRange", fmt.Sprintf("%v", request.outPositionRange))
	if request.beforeText != "" {
		localVarQueryParams.Add("beforeText", fmt.Sprintf("%v", request.beforeText))
	}
	if request.afterText != "" {
		localVarQueryParams.Add("afterText", fmt.Sprintf("%v", request.afterText))
	}
	if request.beforePosition != nil {
		localVarQueryParams.Add("beforePosition", fmt.Sprintf("%v", *request.beforePosition))
	}
	if request.afterPosition != nil {
		localVarQueryParams.Add("afterPosition", fmt.Sprintf("%v", *request.afterPosition))
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

func (request *ExtractTextRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ExtractTextRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
