package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type DeleteFolderRequest struct {
	path string

	recursive   *bool
	storageName string

	extraQueryParameters map[string]string
}

func NewDeleteFolderRequest(path string, opts ...Option) *DeleteFolderRequest {
	req := &DeleteFolderRequest{
		path: path,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["recursive"].(*bool); ok {
		req.recursive = val
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

func (request *DeleteFolderRequest) Validate() error {
	if request.path == "" {
		return fmt.Errorf("required request parameter %q is missing", "path")
	}

	return nil
}

func (request *DeleteFolderRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *DeleteFolderRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *DeleteFolderRequest) GetMethod() string {
	return "DELETE"
}

func (request *DeleteFolderRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *DeleteFolderRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(fmt.Sprintf("%v", request.path)), -1)
	return localVarPath
}

func (request *DeleteFolderRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.recursive != nil {
		localVarQueryParams.Add("recursive", fmt.Sprintf("%v", *request.recursive))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *DeleteFolderRequest) GetJSONBody() interface{} {
	return nil
}

func (request *DeleteFolderRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
