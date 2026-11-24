package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorkbookCalculateFormulaRequest struct {
	name string

	folder      string
	ignoreError *bool
	options     *models.CalculationOptions
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorkbookCalculateFormulaRequest(name string, opts ...Option) *PostWorkbookCalculateFormulaRequest {
	req := &PostWorkbookCalculateFormulaRequest{
		name: name,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["ignoreError"].(*bool); ok {
		req.ignoreError = val
	}
	if val, ok := cfg.Params["options"].(*models.CalculationOptions); ok {
		req.options = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
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

func (request *PostWorkbookCalculateFormulaRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostWorkbookCalculateFormulaRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookCalculateFormulaRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookCalculateFormulaRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookCalculateFormulaRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookCalculateFormulaRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/calculateformula"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbookCalculateFormulaRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.ignoreError != nil {
		localVarQueryParams.Add("ignoreError", fmt.Sprintf("%v", *request.ignoreError))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostWorkbookCalculateFormulaRequest) GetJSONBody() interface{} {
	return &request.options
}

func (request *PostWorkbookCalculateFormulaRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
