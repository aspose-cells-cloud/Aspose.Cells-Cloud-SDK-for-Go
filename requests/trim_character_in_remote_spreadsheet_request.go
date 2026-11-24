package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type TrimCharacterInRemoteSpreadsheetRequest struct {
	name      string
	_range    string
	worksheet string

	folder                  string
	password                string
	region                  string
	removeAllLineBreaks     *bool
	removeExtraLineBreaks   *bool
	storageName             string
	trimContent             string
	trimLeading             *bool
	trimNonBreakingSpaces   *bool
	trimSpaceBetweenWordTo1 *bool
	trimTrailing            *bool

	extraQueryParameters map[string]string
}

func NewTrimCharacterInRemoteSpreadsheetRequest(name string, _range string, worksheet string, opts ...Option) *TrimCharacterInRemoteSpreadsheetRequest {
	req := &TrimCharacterInRemoteSpreadsheetRequest{
		name:      name,
		_range:    _range,
		worksheet: worksheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
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
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
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

func (request *TrimCharacterInRemoteSpreadsheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/range/{range}/content/trim"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", url.PathEscape(fmt.Sprintf("%v", request._range)), -1)
	return localVarPath
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
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
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
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

func (request *TrimCharacterInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *TrimCharacterInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
