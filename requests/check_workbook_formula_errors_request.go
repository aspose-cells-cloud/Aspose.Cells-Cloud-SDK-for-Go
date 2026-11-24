package requests

import (
	"fmt"
	"net/url"

	"asposecellscloud/models"
)

type CheckWorkbookFormulaErrorsRequest struct {
	formulaErrorOptions *models.CheckFormulaErrorOptions

	extraQueryParameters map[string]string
}

func NewCheckWorkbookFormulaErrorsRequest(formulaErrorOptions *models.CheckFormulaErrorOptions, opts ...Option) *CheckWorkbookFormulaErrorsRequest {
	req := &CheckWorkbookFormulaErrorsRequest{
		formulaErrorOptions: formulaErrorOptions,
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

func (request *CheckWorkbookFormulaErrorsRequest) Validate() error {
	if request.formulaErrorOptions == nil {
		return fmt.Errorf("required request parameter %q is missing", "formulaErrorOptions")
	}

	return nil
}

func (request *CheckWorkbookFormulaErrorsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CheckWorkbookFormulaErrorsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CheckWorkbookFormulaErrorsRequest) GetMethod() string {
	return "POST"
}

func (request *CheckWorkbookFormulaErrorsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *CheckWorkbookFormulaErrorsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/checkformulaerrors"
	return localVarPath
}

func (request *CheckWorkbookFormulaErrorsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *CheckWorkbookFormulaErrorsRequest) GetJSONBody() interface{} {
	return &request.formulaErrorOptions
}

func (request *CheckWorkbookFormulaErrorsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
