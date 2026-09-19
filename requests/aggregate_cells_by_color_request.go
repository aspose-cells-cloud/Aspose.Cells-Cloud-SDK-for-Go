package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type AggregateCellsByColorRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	colorPosition string
	operation     string
	password      string
	_range        string
	region        string
	worksheet     string

	extraQueryParameters map[string]string
}

func NewAggregateCellsByColorRequest(Spreadsheet string, opts ...Option) *AggregateCellsByColorRequest {
	req := &AggregateCellsByColorRequest{
		Spreadsheet: Spreadsheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["colorPosition"].(string); ok {
		req.colorPosition = val
	}
	if val, ok := cfg.Params["operation"].(string); ok {
		req.operation = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["range"].(string); ok {
		req._range = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["worksheet"].(string); ok {
		req.worksheet = val
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

func (request *AggregateCellsByColorRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *AggregateCellsByColorRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *AggregateCellsByColorRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *AggregateCellsByColorRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *AggregateCellsByColorRequest) GetMethod() string {
	return "PUT"
}

func (request *AggregateCellsByColorRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *AggregateCellsByColorRequest) GetPath() string {
	localVarPath := "/v4.0/cells/calculate/aggergate/color"
	return localVarPath
}

func (request *AggregateCellsByColorRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.worksheet != "" {
		localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	}
	if request._range != "" {
		localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
	}
	if request.operation != "" {
		localVarQueryParams.Add("operation", fmt.Sprintf("%v", request.operation))
	}
	if request.colorPosition != "" {
		localVarQueryParams.Add("colorPosition", fmt.Sprintf("%v", request.colorPosition))
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

func (request *AggregateCellsByColorRequest) GetJSONBody() interface{} {
	return nil
}

func (request *AggregateCellsByColorRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
