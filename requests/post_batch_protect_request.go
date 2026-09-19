package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostBatchProtectRequest struct {
	batchProtectRequest *models.BatchProtectRequest

	extraQueryParameters map[string]string
}

func NewPostBatchProtectRequest(batchProtectRequest *models.BatchProtectRequest, opts ...Option) *PostBatchProtectRequest {
	req := &PostBatchProtectRequest{
		batchProtectRequest: batchProtectRequest,
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

func (request *PostBatchProtectRequest) Validate() error {
	if request.batchProtectRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "batchProtectRequest")
	}

	return nil
}

func (request *PostBatchProtectRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostBatchProtectRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostBatchProtectRequest) GetMethod() string {
	return "POST"
}

func (request *PostBatchProtectRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostBatchProtectRequest) GetPath() string {
	localVarPath := "/v3.0/cells/batch/protect"
	return localVarPath
}

func (request *PostBatchProtectRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostBatchProtectRequest) GetJSONBody() interface{} {
	return &request.batchProtectRequest
}

func (request *PostBatchProtectRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
