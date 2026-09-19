package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostExtractTextRequest struct {
	extractTextOptions *models.ExtractTextOptions

	extraQueryParameters map[string]string
}

func NewPostExtractTextRequest(extractTextOptions *models.ExtractTextOptions, opts ...Option) *PostExtractTextRequest {
	req := &PostExtractTextRequest{
		extractTextOptions: extractTextOptions,
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

func (request *PostExtractTextRequest) Validate() error {
	if request.extractTextOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "extractTextOptions")
	}

	return nil
}

func (request *PostExtractTextRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostExtractTextRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostExtractTextRequest) GetMethod() string {
	return "POST"
}

func (request *PostExtractTextRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostExtractTextRequest) GetPath() string {
	localVarPath := "/v3.0/cells/extracttext"
	return localVarPath
}

func (request *PostExtractTextRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostExtractTextRequest) GetJSONBody() interface{} {
	return &request.extractTextOptions
}

func (request *PostExtractTextRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
