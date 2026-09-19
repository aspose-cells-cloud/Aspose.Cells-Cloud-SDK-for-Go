package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchAllTextItemsInRemoteSpreadsheetRequest struct {
	folder string
	name   string

	password    string
	region      string
	storageName string

	extraQueryParameters map[string]string
}

func NewSearchAllTextItemsInRemoteSpreadsheetRequest(folder string, name string, opts ...Option) *SearchAllTextItemsInRemoteSpreadsheetRequest {
	req := &SearchAllTextItemsInRemoteSpreadsheetRequest{
		folder: folder,
		name:   name,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
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

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) Validate() error {
	if request.folder == "" {
		return fmt.Errorf("required request parameter %q is missing", "folder")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/search/content/all-textitems"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
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

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
