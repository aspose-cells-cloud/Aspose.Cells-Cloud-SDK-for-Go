package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostDataDeduplicationRequest struct {
	dataDeduplicationRequest *models.DataDeduplicationRequest

	extraQueryParameters map[string]string
}

func NewPostDataDeduplicationRequest(dataDeduplicationRequest *models.DataDeduplicationRequest, opts ...Option) *PostDataDeduplicationRequest {
	req := &PostDataDeduplicationRequest{
		dataDeduplicationRequest: dataDeduplicationRequest,
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

func (request *PostDataDeduplicationRequest) Validate() error {
	if request.dataDeduplicationRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "dataDeduplicationRequest")
	}

	return nil
}

func (request *PostDataDeduplicationRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostDataDeduplicationRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostDataDeduplicationRequest) GetMethod() string {
	return "POST"
}

func (request *PostDataDeduplicationRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostDataDeduplicationRequest) GetPath() string {
	localVarPath := "/v3.0/cells/datadeduplication"
	return localVarPath
}

func (request *PostDataDeduplicationRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostDataDeduplicationRequest) GetJSONBody() interface{} {
	return &request.dataDeduplicationRequest
}

func (request *PostDataDeduplicationRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
