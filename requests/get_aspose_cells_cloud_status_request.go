package requests

import (
	"net/url"
)

type GetAsposeCellsCloudStatusRequest struct {
	extraQueryParameters map[string]string
}

func NewGetAsposeCellsCloudStatusRequest(opts ...Option) *GetAsposeCellsCloudStatusRequest {
	req := &GetAsposeCellsCloudStatusRequest{}
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

func (request *GetAsposeCellsCloudStatusRequest) Validate() error {
	return nil
}

func (request *GetAsposeCellsCloudStatusRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetAsposeCellsCloudStatusRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetAsposeCellsCloudStatusRequest) GetMethod() string {
	return "GET"
}

func (request *GetAsposeCellsCloudStatusRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetAsposeCellsCloudStatusRequest) GetPath() string {
	localVarPath := "/v4.0/cells"
	return localVarPath
}

func (request *GetAsposeCellsCloudStatusRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *GetAsposeCellsCloudStatusRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetAsposeCellsCloudStatusRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
