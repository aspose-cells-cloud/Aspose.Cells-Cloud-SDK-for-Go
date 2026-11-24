package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostAddTextContentRequest struct {
	addTextOptions *models.AddTextOptions

	extraQueryParameters map[string]string
}

func NewPostAddTextContentRequest(addTextOptions *models.AddTextOptions, opts ...Option) *PostAddTextContentRequest {
	req := &PostAddTextContentRequest{
		addTextOptions: addTextOptions,
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

func (request *PostAddTextContentRequest) Validate() error {
	if request.addTextOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "addTextOptions")
	}

	return nil
}

func (request *PostAddTextContentRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostAddTextContentRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostAddTextContentRequest) GetMethod() string {
	return "POST"
}

func (request *PostAddTextContentRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostAddTextContentRequest) GetPath() string {
	localVarPath := "/v3.0/cells/addtext"
	return localVarPath
}

func (request *PostAddTextContentRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostAddTextContentRequest) GetJSONBody() interface{} {
	return &request.addTextOptions
}

func (request *PostAddTextContentRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
