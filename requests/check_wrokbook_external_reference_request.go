package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type CheckWrokbookExternalReferenceRequest struct {
	checkExternalReferenceOptions *models.CheckExternalReferenceOptions

	extraQueryParameters map[string]string
}

func NewCheckWrokbookExternalReferenceRequest(checkExternalReferenceOptions *models.CheckExternalReferenceOptions, opts ...Option) *CheckWrokbookExternalReferenceRequest {
	req := &CheckWrokbookExternalReferenceRequest{
		checkExternalReferenceOptions: checkExternalReferenceOptions,
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

func (request *CheckWrokbookExternalReferenceRequest) Validate() error {
	if request.checkExternalReferenceOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "checkExternalReferenceOptions")
	}

	return nil
}

func (request *CheckWrokbookExternalReferenceRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CheckWrokbookExternalReferenceRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CheckWrokbookExternalReferenceRequest) GetMethod() string {
	return "POST"
}

func (request *CheckWrokbookExternalReferenceRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *CheckWrokbookExternalReferenceRequest) GetPath() string {
	localVarPath := "/v3.0/cells/checkexternalreference"
	return localVarPath
}

func (request *CheckWrokbookExternalReferenceRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *CheckWrokbookExternalReferenceRequest) GetJSONBody() interface{} {
	return &request.checkExternalReferenceOptions
}

func (request *CheckWrokbookExternalReferenceRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
