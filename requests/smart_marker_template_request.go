package requests

import (
	"fmt"
	"net/url"
)

type SmartMarkerTemplateRequest struct {
	password string
	region   string

	extraQueryParameters map[string]string
}

func NewSmartMarkerTemplateRequest(opts ...Option) *SmartMarkerTemplateRequest {
	req := &SmartMarkerTemplateRequest{}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
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

func (request *SmartMarkerTemplateRequest) Validate() error {
	return nil
}

func (request *SmartMarkerTemplateRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SmartMarkerTemplateRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SmartMarkerTemplateRequest) GetMethod() string {
	return "PUT"
}

func (request *SmartMarkerTemplateRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *SmartMarkerTemplateRequest) GetPath() string {
	localVarPath := "/v4.0/cells/report/smart/template"
	return localVarPath
}

func (request *SmartMarkerTemplateRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *SmartMarkerTemplateRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SmartMarkerTemplateRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
