package requests

import (
	"fmt"
	"net/url"
)

type GetDiscUsageRequest struct {
	storageName string

	extraQueryParameters map[string]string
}

func NewGetDiscUsageRequest(opts ...Option) *GetDiscUsageRequest {
	req := &GetDiscUsageRequest{}
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

func (request *GetDiscUsageRequest) Validate() error {
	return nil
}

func (request *GetDiscUsageRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetDiscUsageRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetDiscUsageRequest) GetMethod() string {
	return "GET"
}

func (request *GetDiscUsageRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetDiscUsageRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/disc"
	return localVarPath
}

func (request *GetDiscUsageRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *GetDiscUsageRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetDiscUsageRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
