package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type GetWorksheetsWithLocalSpreadsheetRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	password string
	region   string

	extraQueryParameters map[string]string
}

func NewGetWorksheetsWithLocalSpreadsheetRequest(Spreadsheet string, opts ...Option) *GetWorksheetsWithLocalSpreadsheetRequest {
	req := &GetWorksheetsWithLocalSpreadsheetRequest{
		Spreadsheet: Spreadsheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
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

func (request *GetWorksheetsWithLocalSpreadsheetRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/spreadsheet/worksheets"
	return localVarPath
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
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

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
