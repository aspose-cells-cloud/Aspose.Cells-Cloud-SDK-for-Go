package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type GetDocumentPropertyRequest struct {
	name         string
	propertyName string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewGetDocumentPropertyRequest(name string, propertyName string, opts ...Option) *GetDocumentPropertyRequest {
	req := &GetDocumentPropertyRequest{
		name:         name,
		propertyName: propertyName,
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

func (request *GetDocumentPropertyRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.propertyName == "" {
		return fmt.Errorf("required request parameter %q is missing", "propertyName")
	}

	return nil
}

func (request *GetDocumentPropertyRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetDocumentPropertyRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetDocumentPropertyRequest) GetMethod() string {
	return "GET"
}

func (request *GetDocumentPropertyRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetDocumentPropertyRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(fmt.Sprintf("%v", request.propertyName)), -1)
	return localVarPath
}

func (request *GetDocumentPropertyRequest) GetQueryParameters() url.Values {
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

func (request *GetDocumentPropertyRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetDocumentPropertyRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
