package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type ImportXMLDataIntoSpreadsheetRequest struct {
	datafile        string
	datafileData    []byte
	datafileName    string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	startcell       string
	worksheet       string

	fontsLocation  string
	insert         *bool
	outPath        string
	outStorageName string
	password       string
	region         string

	extraQueryParameters map[string]string
}

func NewImportXMLDataIntoSpreadsheetRequest(datafile string, Spreadsheet string, startcell string, worksheet string, opts ...Option) *ImportXMLDataIntoSpreadsheetRequest {
	req := &ImportXMLDataIntoSpreadsheetRequest{
		datafile:    datafile,
		Spreadsheet: Spreadsheet,
		startcell:   startcell,
		worksheet:   worksheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["fontsLocation"].(string); ok {
		req.fontsLocation = val
	}
	if val, ok := cfg.Params["insert"].(*bool); ok {
		req.insert = val
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

func (request *ImportXMLDataIntoSpreadsheetRequest) Validate() error {
	if request.datafileData == nil && request.datafile == "" {
		return fmt.Errorf("required request parameter %q is missing", "datafile")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.startcell == "" {
		return fmt.Errorf("required request parameter %q is missing", "startcell")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *ImportXMLDataIntoSpreadsheetRequest) SetDatafileBytes(data []byte, name string) {
	if name == "" {
		name = "datafile"
	}
	request.datafileData = data
	request.datafileName = name
}

func (request *ImportXMLDataIntoSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *ImportXMLDataIntoSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ImportXMLDataIntoSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ImportXMLDataIntoSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *ImportXMLDataIntoSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *ImportXMLDataIntoSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/import/data/xml"
	return localVarPath
}

func (request *ImportXMLDataIntoSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	localVarQueryParams.Add("startcell", fmt.Sprintf("%v", request.startcell))
	if request.insert != nil {
		localVarQueryParams.Add("insert", fmt.Sprintf("%v", *request.insert))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.fontsLocation != "" {
		localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
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

func (request *ImportXMLDataIntoSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ImportXMLDataIntoSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.datafileData != nil {
		localVarFormParams[request.datafileName] = request.datafileData
	} else if request.datafile != "" {
		localVarFormParams["@"+filepath.Base(request.datafile)] = request.datafile
	}
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
