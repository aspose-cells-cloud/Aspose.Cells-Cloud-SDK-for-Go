package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type ReplaceContentInRemoteWorksheetRequest struct {
	name        string
	replaceText string
	searchText  string
	worksheet   string

	folder      string
	password    string
	region      string
	storageName string

	extraQueryParameters map[string]string
}

func NewReplaceContentInRemoteWorksheetRequest(name string, replaceText string, searchText string, worksheet string, opts ...Option) *ReplaceContentInRemoteWorksheetRequest {
	req := &ReplaceContentInRemoteWorksheetRequest{
		name:        name,
		replaceText: replaceText,
		searchText:  searchText,
		worksheet:   worksheet,
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

func (request *ReplaceContentInRemoteWorksheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.replaceText == "" {
		return fmt.Errorf("required request parameter %q is missing", "replaceText")
	}

	if request.searchText == "" {
		return fmt.Errorf("required request parameter %q is missing", "searchText")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *ReplaceContentInRemoteWorksheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ReplaceContentInRemoteWorksheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ReplaceContentInRemoteWorksheetRequest) GetMethod() string {
	return "PUT"
}

func (request *ReplaceContentInRemoteWorksheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *ReplaceContentInRemoteWorksheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/replace/content"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	return localVarPath
}

func (request *ReplaceContentInRemoteWorksheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("searchText", fmt.Sprintf("%v", request.searchText))
	localVarQueryParams.Add("replaceText", fmt.Sprintf("%v", request.replaceText))
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

func (request *ReplaceContentInRemoteWorksheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ReplaceContentInRemoteWorksheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
