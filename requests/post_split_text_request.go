package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostSplitTextRequest struct {
	splitTextOptions *models.SplitTextOptions

	extraQueryParameters map[string]string
}

func NewPostSplitTextRequest(splitTextOptions *models.SplitTextOptions, opts ...Option) *PostSplitTextRequest {
	req := &PostSplitTextRequest{
		splitTextOptions: splitTextOptions,
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

func (request *PostSplitTextRequest) Validate() error {
	if request.splitTextOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "splitTextOptions")
	}

	return nil
}

func (request *PostSplitTextRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostSplitTextRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostSplitTextRequest) GetMethod() string {
	return "POST"
}

func (request *PostSplitTextRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostSplitTextRequest) GetPath() string {
	localVarPath := "/v3.0/cells/splittext"
	return localVarPath
}

func (request *PostSplitTextRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostSplitTextRequest) GetJSONBody() interface{} {
	return &request.splitTextOptions
}

func (request *PostSplitTextRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
