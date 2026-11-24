package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type ImportDataIntoSpreadsheetRequest struct {
	datafile        string
	datafileData    []byte
	datafileName    string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	startcell       string
	worksheet       string

	convertNumericData *bool
	fontsLocation      string
	insert             *bool
	outPath            string
	outStorageName     string
	password           string
	region             string
	splitter           string

	extraQueryParameters map[string]string
}

func NewImportDataIntoSpreadsheetRequest(datafile string, Spreadsheet string, startcell string, worksheet string, opts ...Option) *ImportDataIntoSpreadsheetRequest {
	req := &ImportDataIntoSpreadsheetRequest{
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

	if val, ok := cfg.Params["convertNumericData"].(*bool); ok {
		req.convertNumericData = val
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
	if val, ok := cfg.Params["splitter"].(string); ok {
		req.splitter = val
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

func (request *ImportDataIntoSpreadsheetRequest) Validate() error {
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

func (request *ImportDataIntoSpreadsheetRequest) SetDatafileBytes(data []byte, name string) {
	if name == "" {
		name = "datafile"
	}
	request.datafileData = data
	request.datafileName = name
}

func (request *ImportDataIntoSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *ImportDataIntoSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ImportDataIntoSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ImportDataIntoSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *ImportDataIntoSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *ImportDataIntoSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/import/data"
	return localVarPath
}

func (request *ImportDataIntoSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	localVarQueryParams.Add("startcell", fmt.Sprintf("%v", request.startcell))
	if request.insert != nil {
		localVarQueryParams.Add("insert", fmt.Sprintf("%v", *request.insert))
	}
	if request.convertNumericData != nil {
		localVarQueryParams.Add("convertNumericData", fmt.Sprintf("%v", *request.convertNumericData))
	}
	if request.splitter != "" {
		localVarQueryParams.Add("splitter", fmt.Sprintf("%v", request.splitter))
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

func (request *ImportDataIntoSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ImportDataIntoSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
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
