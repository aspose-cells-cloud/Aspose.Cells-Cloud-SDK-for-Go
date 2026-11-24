package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostConvertWorksheetToImageRequest struct {
	convertWorksheetOptions *models.ConvertWorksheetOptions

	FontsLocation string

	extraQueryParameters map[string]string
}

func NewPostConvertWorksheetToImageRequest(convertWorksheetOptions *models.ConvertWorksheetOptions, opts ...Option) *PostConvertWorksheetToImageRequest {
	req := &PostConvertWorksheetToImageRequest{
		convertWorksheetOptions: convertWorksheetOptions,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["FontsLocation"].(string); ok {
		req.FontsLocation = val
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

func (request *PostConvertWorksheetToImageRequest) Validate() error {
	if request.convertWorksheetOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "convertWorksheetOptions")
	}

	return nil
}

func (request *PostConvertWorksheetToImageRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostConvertWorksheetToImageRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostConvertWorksheetToImageRequest) GetMethod() string {
	return "POST"
}

func (request *PostConvertWorksheetToImageRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostConvertWorksheetToImageRequest) GetPath() string {
	localVarPath := "/v3.0/cells/convertWorksheetToImage"
	return localVarPath
}

func (request *PostConvertWorksheetToImageRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.FontsLocation != "" {
		localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostConvertWorksheetToImageRequest) GetJSONBody() interface{} {
	return &request.convertWorksheetOptions
}

func (request *PostConvertWorksheetToImageRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
