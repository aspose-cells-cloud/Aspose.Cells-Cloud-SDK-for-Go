package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostDataTransformationRequest struct {
	dataTransformationRequest *models.DataTransformationRequest

	extraQueryParameters map[string]string
}

func NewPostDataTransformationRequest(dataTransformationRequest *models.DataTransformationRequest, opts ...Option) *PostDataTransformationRequest {
	req := &PostDataTransformationRequest{
		dataTransformationRequest: dataTransformationRequest,
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

func (request *PostDataTransformationRequest) Validate() error {
	if request.dataTransformationRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "dataTransformationRequest")
	}

	return nil
}

func (request *PostDataTransformationRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostDataTransformationRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostDataTransformationRequest) GetMethod() string {
	return "POST"
}

func (request *PostDataTransformationRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostDataTransformationRequest) GetPath() string {
	localVarPath := "/v3.0/cells/datatransformation"
	return localVarPath
}

func (request *PostDataTransformationRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostDataTransformationRequest) GetJSONBody() interface{} {
	return &request.dataTransformationRequest
}

func (request *PostDataTransformationRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
