package requests

import (
	"fmt"
	"net/url"
)

type CreateSpreadsheetRequest struct {
	format         string
	outPath        string
	outStorageName string
	password       string
	region         string
	template       string

	extraQueryParameters map[string]string
}

func NewCreateSpreadsheetRequest(opts ...Option) *CreateSpreadsheetRequest {
	req := &CreateSpreadsheetRequest{}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["format"].(string); ok {
		req.format = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["template"].(string); ok {
		req.template = val
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

func (request *CreateSpreadsheetRequest) Validate() error {
	return nil
}

func (request *CreateSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CreateSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CreateSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *CreateSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *CreateSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/spreadsheet/create"
	return localVarPath
}

func (request *CreateSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.format != "" {
		localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
	}
	if request.template != "" {
		localVarQueryParams.Add("template", fmt.Sprintf("%v", request.template))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *CreateSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *CreateSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
