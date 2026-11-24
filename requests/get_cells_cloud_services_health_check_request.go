package requests

import (
	"net/url"
)

type GetCellsCloudServicesHealthCheckRequest struct {
	extraQueryParameters map[string]string
}

func NewGetCellsCloudServicesHealthCheckRequest(opts ...Option) *GetCellsCloudServicesHealthCheckRequest {
	req := &GetCellsCloudServicesHealthCheckRequest{}
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

func (request *GetCellsCloudServicesHealthCheckRequest) Validate() error {
	return nil
}

func (request *GetCellsCloudServicesHealthCheckRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetCellsCloudServicesHealthCheckRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetMethod() string {
	return "GET"
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetPath() string {
	localVarPath := "/v3.0/cells"
	return localVarPath
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
