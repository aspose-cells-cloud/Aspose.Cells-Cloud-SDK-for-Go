package requests

import (
	"net/url"
)

type GetPublicKeyRequest struct {
	extraQueryParameters map[string]string
}

func NewGetPublicKeyRequest(opts ...Option) *GetPublicKeyRequest {
	req := &GetPublicKeyRequest{}
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

func (request *GetPublicKeyRequest) Validate() error {
	return nil
}

func (request *GetPublicKeyRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetPublicKeyRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetPublicKeyRequest) GetMethod() string {
	return "GET"
}

func (request *GetPublicKeyRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetPublicKeyRequest) GetPath() string {
	localVarPath := "/v4.0/cells/publickey"
	return localVarPath
}

func (request *GetPublicKeyRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *GetPublicKeyRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetPublicKeyRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
