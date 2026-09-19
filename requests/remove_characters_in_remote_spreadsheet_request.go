package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type RemoveCharactersInRemoteSpreadsheetRequest struct {
	name      string
	_range    string
	worksheet string

	caseSensitive     *bool
	characterSets     string
	folder            string
	password          string
	region            string
	removeCustomValue string
	removeTextMethod  string
	storageName       string

	extraQueryParameters map[string]string
}

func NewRemoveCharactersInRemoteSpreadsheetRequest(name string, _range string, worksheet string, opts ...Option) *RemoveCharactersInRemoteSpreadsheetRequest {
	req := &RemoveCharactersInRemoteSpreadsheetRequest{
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

	if val, ok := cfg.Params["caseSensitive"].(*bool); ok {
		req.caseSensitive = val
	}
	if val, ok := cfg.Params["characterSets"].(string); ok {
		req.characterSets = val
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
	if val, ok := cfg.Params["removeCustomValue"].(string); ok {
		req.removeCustomValue = val
	}
	if val, ok := cfg.Params["removeTextMethod"].(string); ok {
		req.removeTextMethod = val
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

func (request *RemoveCharactersInRemoteSpreadsheetRequest) Validate() error {
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

func (request *RemoveCharactersInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/range/{range}/content/remove/characters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", url.PathEscape(fmt.Sprintf("%v", request._range)), -1)
	return localVarPath
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.removeTextMethod != "" {
		localVarQueryParams.Add("removeTextMethod", fmt.Sprintf("%v", request.removeTextMethod))
	}
	if request.characterSets != "" {
		localVarQueryParams.Add("characterSets", fmt.Sprintf("%v", request.characterSets))
	}
	if request.removeCustomValue != "" {
		localVarQueryParams.Add("removeCustomValue", fmt.Sprintf("%v", request.removeCustomValue))
	}
	if request.caseSensitive != nil {
		localVarQueryParams.Add("caseSensitive", fmt.Sprintf("%v", *request.caseSensitive))
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

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
