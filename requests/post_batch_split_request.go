package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostBatchSplitRequest struct {
	batchSplitRequest *models.BatchSplitRequest

	extraQueryParameters map[string]string
}

func NewPostBatchSplitRequest(batchSplitRequest *models.BatchSplitRequest, opts ...Option) *PostBatchSplitRequest {
	req := &PostBatchSplitRequest{
		batchSplitRequest: batchSplitRequest,
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

func (request *PostBatchSplitRequest) Validate() error {
	if request.batchSplitRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "batchSplitRequest")
	}

	return nil
}

func (request *PostBatchSplitRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostBatchSplitRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostBatchSplitRequest) GetMethod() string {
	return "POST"
}

func (request *PostBatchSplitRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostBatchSplitRequest) GetPath() string {
	localVarPath := "/v3.0/cells/batch/split"
	return localVarPath
}

func (request *PostBatchSplitRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostBatchSplitRequest) GetJSONBody() interface{} {
	return &request.batchSplitRequest
}

func (request *PostBatchSplitRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
