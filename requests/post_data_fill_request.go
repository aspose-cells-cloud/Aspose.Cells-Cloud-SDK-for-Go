package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostDataFillRequest struct {
	dataFillRequest *models.DataFillRequest

	extraQueryParameters map[string]string
}

func NewPostDataFillRequest(dataFillRequest *models.DataFillRequest, opts ...Option) *PostDataFillRequest {
	req := &PostDataFillRequest{
		dataFillRequest: dataFillRequest,
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

func (request *PostDataFillRequest) Validate() error {
	if request.dataFillRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "dataFillRequest")
	}

	return nil
}

func (request *PostDataFillRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostDataFillRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostDataFillRequest) GetMethod() string {
	return "POST"
}

func (request *PostDataFillRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostDataFillRequest) GetPath() string {
	localVarPath := "/v3.0/cells/datafill"
	return localVarPath
}

func (request *PostDataFillRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostDataFillRequest) GetJSONBody() interface{} {
	return &request.dataFillRequest
}

func (request *PostDataFillRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
