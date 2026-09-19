package requests

import (
	"net/url"
)

type CheckCloudServiceHealthRequest struct {
	extraQueryParameters map[string]string
}

func NewCheckCloudServiceHealthRequest(opts ...Option) *CheckCloudServiceHealthRequest {
	req := &CheckCloudServiceHealthRequest{}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
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

func (request *CheckCloudServiceHealthRequest) Validate() error {
	return nil
}

func (request *CheckCloudServiceHealthRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CheckCloudServiceHealthRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CheckCloudServiceHealthRequest) GetMethod() string {
	return "GET"
}

func (request *CheckCloudServiceHealthRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *CheckCloudServiceHealthRequest) GetPath() string {
	localVarPath := "/v4.0/cells/status/check"
	return localVarPath
}

func (request *CheckCloudServiceHealthRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *CheckCloudServiceHealthRequest) GetJSONBody() interface{} {
	return nil
}

func (request *CheckCloudServiceHealthRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
