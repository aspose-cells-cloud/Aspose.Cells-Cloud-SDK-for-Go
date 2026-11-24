package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostDataCleansingRequest struct {
	dataCleansingRequest *models.DataCleansingRequest

	extraQueryParameters map[string]string
}

func NewPostDataCleansingRequest(dataCleansingRequest *models.DataCleansingRequest, opts ...Option) *PostDataCleansingRequest {
	req := &PostDataCleansingRequest{
		dataCleansingRequest: dataCleansingRequest,
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

func (request *PostDataCleansingRequest) Validate() error {
	if request.dataCleansingRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "dataCleansingRequest")
	}

	return nil
}

func (request *PostDataCleansingRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostDataCleansingRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostDataCleansingRequest) GetMethod() string {
	return "POST"
}

func (request *PostDataCleansingRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostDataCleansingRequest) GetPath() string {
	localVarPath := "/v3.0/cells/datacleansing"
	return localVarPath
}

func (request *PostDataCleansingRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostDataCleansingRequest) GetJSONBody() interface{} {
	return &request.dataCleansingRequest
}

func (request *PostDataCleansingRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
