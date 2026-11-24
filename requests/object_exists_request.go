package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type ObjectExistsRequest struct {
	path string

	storageName string
	versionId   string

	extraQueryParameters map[string]string
}

func NewObjectExistsRequest(path string, opts ...Option) *ObjectExistsRequest {
	req := &ObjectExistsRequest{
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
	if val, ok := cfg.Params["versionId"].(string); ok {
		req.versionId = val
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

func (request *ObjectExistsRequest) Validate() error {
	if request.path == "" {
		return fmt.Errorf("required request parameter %q is missing", "path")
	}

	return nil
}

func (request *ObjectExistsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ObjectExistsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ObjectExistsRequest) GetMethod() string {
	return "GET"
}

func (request *ObjectExistsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *ObjectExistsRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/exist/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(fmt.Sprintf("%v", request.path)), -1)
	return localVarPath
}

func (request *ObjectExistsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.versionId != "" {
		localVarQueryParams.Add("versionId", fmt.Sprintf("%v", request.versionId))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *ObjectExistsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ObjectExistsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
