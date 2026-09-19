package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostTrimContentRequest struct {
	trimContentOptions *models.TrimContentOptions

	extraQueryParameters map[string]string
}

func NewPostTrimContentRequest(trimContentOptions *models.TrimContentOptions, opts ...Option) *PostTrimContentRequest {
	req := &PostTrimContentRequest{
		trimContentOptions: trimContentOptions,
	}
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

func (request *PostTrimContentRequest) Validate() error {
	if request.trimContentOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "trimContentOptions")
	}

	return nil
}

func (request *PostTrimContentRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostTrimContentRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostTrimContentRequest) GetMethod() string {
	return "POST"
}

func (request *PostTrimContentRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostTrimContentRequest) GetPath() string {
	localVarPath := "/v3.0/cells/trimcontent"
	return localVarPath
}

func (request *PostTrimContentRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostTrimContentRequest) GetJSONBody() interface{} {
	return &request.trimContentOptions
}

func (request *PostTrimContentRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
