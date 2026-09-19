package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostSpecifyWordsCountRequest struct {
	specifyWordsCountOptions *models.SpecifyWordsCountOptions

	extraQueryParameters map[string]string
}

func NewPostSpecifyWordsCountRequest(specifyWordsCountOptions *models.SpecifyWordsCountOptions, opts ...Option) *PostSpecifyWordsCountRequest {
	req := &PostSpecifyWordsCountRequest{
		specifyWordsCountOptions: specifyWordsCountOptions,
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

func (request *PostSpecifyWordsCountRequest) Validate() error {
	if request.specifyWordsCountOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "specifyWordsCountOptions")
	}

	return nil
}

func (request *PostSpecifyWordsCountRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostSpecifyWordsCountRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostSpecifyWordsCountRequest) GetMethod() string {
	return "POST"
}

func (request *PostSpecifyWordsCountRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostSpecifyWordsCountRequest) GetPath() string {
	localVarPath := "/v3.0/cells/specifywordscount"
	return localVarPath
}

func (request *PostSpecifyWordsCountRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostSpecifyWordsCountRequest) GetJSONBody() interface{} {
	return &request.specifyWordsCountOptions
}

func (request *PostSpecifyWordsCountRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
