package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type CreateFolderRequest struct {
	path string

	storageName string

	extraQueryParameters map[string]string
}

func NewCreateFolderRequest(path string, opts ...Option) *CreateFolderRequest {
	req := &CreateFolderRequest{
		path: path,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
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

func (request *CreateFolderRequest) Validate() error {
	if request.path == "" {
		return fmt.Errorf("required request parameter %q is missing", "path")
	}

	return nil
}

func (request *CreateFolderRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CreateFolderRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CreateFolderRequest) GetMethod() string {
	return "PUT"
}

func (request *CreateFolderRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *CreateFolderRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(fmt.Sprintf("%v", request.path)), -1)
	return localVarPath
}

func (request *CreateFolderRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *CreateFolderRequest) GetJSONBody() interface{} {
	return nil
}

func (request *CreateFolderRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
