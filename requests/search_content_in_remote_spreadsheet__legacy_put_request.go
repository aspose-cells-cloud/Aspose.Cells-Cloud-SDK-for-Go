package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchContentInRemoteSpreadsheet_LegacyPutRequest struct {
	name       string
	searchText string

	folder       string
	ignoringCase *bool
	password     string
	region       string
	storageName  string

	extraQueryParameters map[string]string
}

func NewSearchContentInRemoteSpreadsheet_LegacyPutRequest(name string, searchText string, opts ...Option) *SearchContentInRemoteSpreadsheet_LegacyPutRequest {
	req := &SearchContentInRemoteSpreadsheet_LegacyPutRequest{
		name:       name,
		searchText: searchText,
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
	if val, ok := cfg.Params["ignoringCase"].(*bool); ok {
		req.ignoringCase = val
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

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.searchText == "" {
		return fmt.Errorf("required request parameter %q is missing", "searchText")
	}

	return nil
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) GetMethod() string {
	return "PUT"
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/search/content"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("searchText", fmt.Sprintf("%v", request.searchText))
	if request.ignoringCase != nil {
		localVarQueryParams.Add("ignoringCase", fmt.Sprintf("%v", *request.ignoringCase))
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

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SearchContentInRemoteSpreadsheet_LegacyPutRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
