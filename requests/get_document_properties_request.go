package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type GetDocumentPropertiesRequest struct {
	name string

	folder      string
	storageName string
	_type       string

	extraQueryParameters map[string]string
}

func NewGetDocumentPropertiesRequest(name string, opts ...Option) *GetDocumentPropertiesRequest {
	req := &GetDocumentPropertiesRequest{
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
	if val, ok := cfg.Params["type"].(string); ok {
		req._type = val
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

func (request *GetDocumentPropertiesRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *GetDocumentPropertiesRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetDocumentPropertiesRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetDocumentPropertiesRequest) GetMethod() string {
	return "GET"
}

func (request *GetDocumentPropertiesRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetDocumentPropertiesRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/documentproperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *GetDocumentPropertiesRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request._type != "" {
		localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
	}
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

func (request *GetDocumentPropertiesRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetDocumentPropertiesRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
