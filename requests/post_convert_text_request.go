package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostConvertTextRequest struct {
	convertTextOptions *models.ConvertTextOptions

	extraQueryParameters map[string]string
}

func NewPostConvertTextRequest(convertTextOptions *models.ConvertTextOptions, opts ...Option) *PostConvertTextRequest {
	req := &PostConvertTextRequest{
		convertTextOptions: convertTextOptions,
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

func (request *PostConvertTextRequest) Validate() error {
	if request.convertTextOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "convertTextOptions")
	}

	return nil
}

func (request *PostConvertTextRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostConvertTextRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostConvertTextRequest) GetMethod() string {
	return "POST"
}

func (request *PostConvertTextRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostConvertTextRequest) GetPath() string {
	localVarPath := "/v3.0/cells/converttext"
	return localVarPath
}

func (request *PostConvertTextRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostConvertTextRequest) GetJSONBody() interface{} {
	return &request.convertTextOptions
}

func (request *PostConvertTextRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
