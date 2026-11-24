package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostWorkbooksMergeRequest struct {
	mergeWith string
	name      string

	folder            string
	mergedStorageName string
	storageName       string

	extraQueryParameters map[string]string
}

func NewPostWorkbooksMergeRequest(mergeWith string, name string, opts ...Option) *PostWorkbooksMergeRequest {
	req := &PostWorkbooksMergeRequest{
		mergeWith: mergeWith,
		name:      name,
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
	if val, ok := cfg.Params["mergedStorageName"].(string); ok {
		req.mergedStorageName = val
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

func (request *PostWorkbooksMergeRequest) Validate() error {
	if request.mergeWith == "" {
		return fmt.Errorf("required request parameter %q is missing", "mergeWith")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostWorkbooksMergeRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbooksMergeRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbooksMergeRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbooksMergeRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbooksMergeRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/merge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbooksMergeRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("mergeWith", fmt.Sprintf("%v", request.mergeWith))
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.mergedStorageName != "" {
		localVarQueryParams.Add("mergedStorageName", fmt.Sprintf("%v", request.mergedStorageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostWorkbooksMergeRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostWorkbooksMergeRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
