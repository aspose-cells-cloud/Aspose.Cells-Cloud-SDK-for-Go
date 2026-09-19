package requests

import (
	"net/url"
)

type GetCellsCloudServiceStatusRequest struct {
	extraQueryParameters map[string]string
}

func NewGetCellsCloudServiceStatusRequest(opts ...Option) *GetCellsCloudServiceStatusRequest {
	req := &GetCellsCloudServiceStatusRequest{}
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

func (request *GetCellsCloudServiceStatusRequest) Validate() error {
	return nil
}

func (request *GetCellsCloudServiceStatusRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetCellsCloudServiceStatusRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetCellsCloudServiceStatusRequest) GetMethod() string {
	return "GET"
}

func (request *GetCellsCloudServiceStatusRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetCellsCloudServiceStatusRequest) GetPath() string {
	localVarPath := "/v3.0/cells/status/check"
	return localVarPath
}

func (request *GetCellsCloudServiceStatusRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *GetCellsCloudServiceStatusRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetCellsCloudServiceStatusRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
