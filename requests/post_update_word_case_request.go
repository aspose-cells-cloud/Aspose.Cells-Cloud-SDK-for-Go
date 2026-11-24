package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostUpdateWordCaseRequest struct {
	wordCaseOptions *models.WordCaseOptions

	extraQueryParameters map[string]string
}

func NewPostUpdateWordCaseRequest(wordCaseOptions *models.WordCaseOptions, opts ...Option) *PostUpdateWordCaseRequest {
	req := &PostUpdateWordCaseRequest{
		wordCaseOptions: wordCaseOptions,
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

func (request *PostUpdateWordCaseRequest) Validate() error {
	if request.wordCaseOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "wordCaseOptions")
	}

	return nil
}

func (request *PostUpdateWordCaseRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostUpdateWordCaseRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostUpdateWordCaseRequest) GetMethod() string {
	return "POST"
}

func (request *PostUpdateWordCaseRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostUpdateWordCaseRequest) GetPath() string {
	localVarPath := "/v3.0/cells/updatewordcase"
	return localVarPath
}

func (request *PostUpdateWordCaseRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostUpdateWordCaseRequest) GetJSONBody() interface{} {
	return &request.wordCaseOptions
}

func (request *PostUpdateWordCaseRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
