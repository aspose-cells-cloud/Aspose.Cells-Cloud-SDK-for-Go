package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type FlipDataRequest struct {
	cellArea        string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	worksheet       string

	Horizontal     *bool
	outPath        string
	outStorageName string
	password       string
	region         string

	extraQueryParameters map[string]string
}

func NewFlipDataRequest(cellArea string, Spreadsheet string, worksheet string, opts ...Option) *FlipDataRequest {
	req := &FlipDataRequest{
		cellArea:    cellArea,
		Spreadsheet: Spreadsheet,
		worksheet:   worksheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["Horizontal"].(*bool); ok {
		req.Horizontal = val
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

func (request *FlipDataRequest) Validate() error {
	if request.cellArea == "" {
		return fmt.Errorf("required request parameter %q is missing", "cellArea")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *FlipDataRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *FlipDataRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *FlipDataRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *FlipDataRequest) GetMethod() string {
	return "PUT"
}

func (request *FlipDataRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *FlipDataRequest) GetPath() string {
	localVarPath := "/v4.0/cells/flip"
	return localVarPath
}

func (request *FlipDataRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
	if request.Horizontal != nil {
		localVarQueryParams.Add("Horizontal", fmt.Sprintf("%v", *request.Horizontal))
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

func (request *FlipDataRequest) GetJSONBody() interface{} {
	return nil
}

func (request *FlipDataRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
