package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostBatchUnlockRequest struct {
	batchLockRequest *models.BatchLockRequest

	extraQueryParameters map[string]string
}

func NewPostBatchUnlockRequest(batchLockRequest *models.BatchLockRequest, opts ...Option) *PostBatchUnlockRequest {
	req := &PostBatchUnlockRequest{
		batchLockRequest: batchLockRequest,
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

func (request *PostBatchUnlockRequest) Validate() error {
	if request.batchLockRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "batchLockRequest")
	}

	return nil
}

func (request *PostBatchUnlockRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostBatchUnlockRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostBatchUnlockRequest) GetMethod() string {
	return "POST"
}

func (request *PostBatchUnlockRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostBatchUnlockRequest) GetPath() string {
	localVarPath := "/v3.0/cells/batch/unlock"
	return localVarPath
}

func (request *PostBatchUnlockRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostBatchUnlockRequest) GetJSONBody() interface{} {
	return &request.batchLockRequest
}

func (request *PostBatchUnlockRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
