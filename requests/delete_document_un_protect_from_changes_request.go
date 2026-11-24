package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type DeleteDocumentUnProtectFromChangesRequest struct {
	name string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewDeleteDocumentUnProtectFromChangesRequest(name string, opts ...Option) *DeleteDocumentUnProtectFromChangesRequest {
	req := &DeleteDocumentUnProtectFromChangesRequest{
		name: name,
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

func (request *DeleteDocumentUnProtectFromChangesRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *DeleteDocumentUnProtectFromChangesRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *DeleteDocumentUnProtectFromChangesRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetMethod() string {
	return "DELETE"
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/writeProtection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetJSONBody() interface{} {
	return nil
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
