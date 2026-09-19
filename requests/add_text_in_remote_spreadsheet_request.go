package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type AddTextInRemoteSpreadsheetRequest struct {
	name      string
	position  string
	_range    string
	text      string
	worksheet string

	folder         string
	password       string
	region         string
	selectText     string
	skipEmptyCells *bool
	storageName    string

	extraQueryParameters map[string]string
}

func NewAddTextInRemoteSpreadsheetRequest(name string, position string, _range string, text string, worksheet string, opts ...Option) *AddTextInRemoteSpreadsheetRequest {
	req := &AddTextInRemoteSpreadsheetRequest{
		name:      name,
		position:  position,
		_range:    _range,
		text:      text,
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
	if val, ok := cfg.Params["selectText"].(string); ok {
		req.selectText = val
	}
	if val, ok := cfg.Params["skipEmptyCells"].(*bool); ok {
		req.skipEmptyCells = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
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

func (request *AddTextInRemoteSpreadsheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.position == "" {
		return fmt.Errorf("required request parameter %q is missing", "position")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.text == "" {
		return fmt.Errorf("required request parameter %q is missing", "text")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *AddTextInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *AddTextInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *AddTextInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *AddTextInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *AddTextInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/range/{range}/content/add/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", url.PathEscape(fmt.Sprintf("%v", request._range)), -1)
	return localVarPath
}

func (request *AddTextInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
	localVarQueryParams.Add("position", fmt.Sprintf("%v", request.position))
	if request.selectText != "" {
		localVarQueryParams.Add("selectText", fmt.Sprintf("%v", request.selectText))
	}
	if request.skipEmptyCells != nil {
		localVarQueryParams.Add("skipEmptyCells", fmt.Sprintf("%v", *request.skipEmptyCells))
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

func (request *AddTextInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *AddTextInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
