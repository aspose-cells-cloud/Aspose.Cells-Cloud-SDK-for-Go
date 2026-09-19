package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type TrimCharacterRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	outPath                 string
	outStorageName          string
	password                string
	_range                  string
	region                  string
	removeAllLineBreaks     *bool
	removeExtraLineBreaks   *bool
	trimContent             string
	trimLeading             *bool
	trimNonBreakingSpaces   *bool
	trimSpaceBetweenWordTo1 *bool
	trimTrailing            *bool
	worksheet               string

	extraQueryParameters map[string]string
}

func NewTrimCharacterRequest(Spreadsheet string, opts ...Option) *TrimCharacterRequest {
	req := &TrimCharacterRequest{
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
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["range"].(string); ok {
		req._range = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["removeAllLineBreaks"].(*bool); ok {
		req.removeAllLineBreaks = val
	}
	if val, ok := cfg.Params["removeExtraLineBreaks"].(*bool); ok {
		req.removeExtraLineBreaks = val
	}
	if val, ok := cfg.Params["trimContent"].(string); ok {
		req.trimContent = val
	}
	if val, ok := cfg.Params["trimLeading"].(*bool); ok {
		req.trimLeading = val
	}
	if val, ok := cfg.Params["trimNonBreakingSpaces"].(*bool); ok {
		req.trimNonBreakingSpaces = val
	}
	if val, ok := cfg.Params["trimSpaceBetweenWordTo1"].(*bool); ok {
		req.trimSpaceBetweenWordTo1 = val
	}
	if val, ok := cfg.Params["trimTrailing"].(*bool); ok {
		req.trimTrailing = val
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

func (request *TrimCharacterRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *TrimCharacterRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *TrimCharacterRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *TrimCharacterRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *TrimCharacterRequest) GetMethod() string {
	return "PUT"
}

func (request *TrimCharacterRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *TrimCharacterRequest) GetPath() string {
	localVarPath := "/v4.0/cells/content/trim"
	return localVarPath
}

func (request *TrimCharacterRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.trimContent != "" {
		localVarQueryParams.Add("trimContent", fmt.Sprintf("%v", request.trimContent))
	}
	if request.trimLeading != nil {
		localVarQueryParams.Add("trimLeading", fmt.Sprintf("%v", *request.trimLeading))
	}
	if request.trimTrailing != nil {
		localVarQueryParams.Add("trimTrailing", fmt.Sprintf("%v", *request.trimTrailing))
	}
	if request.trimSpaceBetweenWordTo1 != nil {
		localVarQueryParams.Add("trimSpaceBetweenWordTo1", fmt.Sprintf("%v", *request.trimSpaceBetweenWordTo1))
	}
	if request.trimNonBreakingSpaces != nil {
		localVarQueryParams.Add("trimNonBreakingSpaces", fmt.Sprintf("%v", *request.trimNonBreakingSpaces))
	}
	if request.removeExtraLineBreaks != nil {
		localVarQueryParams.Add("removeExtraLineBreaks", fmt.Sprintf("%v", *request.removeExtraLineBreaks))
	}
	if request.removeAllLineBreaks != nil {
		localVarQueryParams.Add("removeAllLineBreaks", fmt.Sprintf("%v", *request.removeAllLineBreaks))
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

func (request *TrimCharacterRequest) GetJSONBody() interface{} {
	return nil
}

func (request *TrimCharacterRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
