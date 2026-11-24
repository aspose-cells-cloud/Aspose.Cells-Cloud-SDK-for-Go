package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type PostAnalyzeExcelRequest struct {
	analyzeExcelRequest *models.AnalyzeExcelRequest

	extraQueryParameters map[string]string
}

func NewPostAnalyzeExcelRequest(analyzeExcelRequest *models.AnalyzeExcelRequest, opts ...Option) *PostAnalyzeExcelRequest {
	req := &PostAnalyzeExcelRequest{
		analyzeExcelRequest: analyzeExcelRequest,
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

func (request *PostAnalyzeExcelRequest) Validate() error {
	if request.analyzeExcelRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "analyzeExcelRequest")
	}

	return nil
}

func (request *PostAnalyzeExcelRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostAnalyzeExcelRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostAnalyzeExcelRequest) GetMethod() string {
	return "POST"
}

func (request *PostAnalyzeExcelRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostAnalyzeExcelRequest) GetPath() string {
	localVarPath := "/v3.0/cells/analyze"
	return localVarPath
}

func (request *PostAnalyzeExcelRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostAnalyzeExcelRequest) GetJSONBody() interface{} {
	return &request.analyzeExcelRequest
}

func (request *PostAnalyzeExcelRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
