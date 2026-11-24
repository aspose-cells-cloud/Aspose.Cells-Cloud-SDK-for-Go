package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type SplitTableRequest struct {
	saveSplitColumn bool
	splitColumnName string
	splitRowNumber  int
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string
	tableName       string
	toMultipleFiles bool
	toNewWorkbook   bool
	worksheet       string

	fontsLocation  string
	outPath        string
	outStorageName string
	password       string
	region         string

	extraQueryParameters map[string]string
}

func NewSplitTableRequest(saveSplitColumn bool, splitColumnName string, splitRowNumber int, Spreadsheet string, tableName string, toMultipleFiles bool, toNewWorkbook bool, worksheet string, opts ...Option) *SplitTableRequest {
	req := &SplitTableRequest{
		saveSplitColumn: saveSplitColumn,
		splitColumnName: splitColumnName,
		splitRowNumber:  splitRowNumber,
		Spreadsheet:     Spreadsheet,
		tableName:       tableName,
		toMultipleFiles: toMultipleFiles,
		toNewWorkbook:   toNewWorkbook,
		worksheet:       worksheet,
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

func (request *SplitTableRequest) Validate() error {
	if request.splitColumnName == "" {
		return fmt.Errorf("required request parameter %q is missing", "splitColumnName")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	if request.tableName == "" {
		return fmt.Errorf("required request parameter %q is missing", "tableName")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *SplitTableRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *SplitTableRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SplitTableRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SplitTableRequest) GetMethod() string {
	return "PUT"
}

func (request *SplitTableRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *SplitTableRequest) GetPath() string {
	localVarPath := "/v4.0/cells/split/table"
	return localVarPath
}

func (request *SplitTableRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	localVarQueryParams.Add("tableName", fmt.Sprintf("%v", request.tableName))
	localVarQueryParams.Add("splitColumnName", fmt.Sprintf("%v", request.splitColumnName))
	localVarQueryParams.Add("saveSplitColumn", fmt.Sprintf("%v", request.saveSplitColumn))
	localVarQueryParams.Add("splitRowNumber", fmt.Sprintf("%v", request.splitRowNumber))
	localVarQueryParams.Add("toNewWorkbook", fmt.Sprintf("%v", request.toNewWorkbook))
	localVarQueryParams.Add("toMultipleFiles", fmt.Sprintf("%v", request.toMultipleFiles))
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

func (request *SplitTableRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SplitTableRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
