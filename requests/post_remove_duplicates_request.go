package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostRemoveDuplicatesRequest struct {
	removeDuplicatesOptions *models.RemoveDuplicatesOptions

	extraQueryParameters map[string]string
}

func NewPostRemoveDuplicatesRequest(removeDuplicatesOptions *models.RemoveDuplicatesOptions, opts ...Option) *PostRemoveDuplicatesRequest {
	req := &PostRemoveDuplicatesRequest{
		removeDuplicatesOptions: removeDuplicatesOptions,
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

func (request *PostRemoveDuplicatesRequest) Validate() error {
	if request.removeDuplicatesOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "removeDuplicatesOptions")
	}

	return nil
}

func (request *PostRemoveDuplicatesRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostRemoveDuplicatesRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostRemoveDuplicatesRequest) GetMethod() string {
	return "POST"
}

func (request *PostRemoveDuplicatesRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostRemoveDuplicatesRequest) GetPath() string {
	localVarPath := "/v3.0/cells/removeduplicates"
	return localVarPath
}

func (request *PostRemoveDuplicatesRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostRemoveDuplicatesRequest) GetJSONBody() interface{} {
	return &request.removeDuplicatesOptions
}

func (request *PostRemoveDuplicatesRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
