package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type GetWorksheetCellsRangeValueRequest struct {
	name      string
	sheetName string

	columnCount *int
	firstColumn *int
	firstRow    *int
	folder      string
	namerange   string
	rowCount    *int
	storageName string

	extraQueryParameters map[string]string
}

func NewGetWorksheetCellsRangeValueRequest(name string, sheetName string, opts ...Option) *GetWorksheetCellsRangeValueRequest {
	req := &GetWorksheetCellsRangeValueRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["columnCount"].(*int); ok {
		req.columnCount = val
	}
	if val, ok := cfg.Params["firstColumn"].(*int); ok {
		req.firstColumn = val
	}
	if val, ok := cfg.Params["firstRow"].(*int); ok {
		req.firstRow = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["namerange"].(string); ok {
		req.namerange = val
	}
	if val, ok := cfg.Params["rowCount"].(*int); ok {
		req.rowCount = val
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

func (request *GetWorksheetCellsRangeValueRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *GetWorksheetCellsRangeValueRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetWorksheetCellsRangeValueRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetWorksheetCellsRangeValueRequest) GetMethod() string {
	return "GET"
}

func (request *GetWorksheetCellsRangeValueRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetWorksheetCellsRangeValueRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/ranges/value"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *GetWorksheetCellsRangeValueRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.namerange != "" {
		localVarQueryParams.Add("namerange", fmt.Sprintf("%v", request.namerange))
	}
	if request.firstRow != nil {
		localVarQueryParams.Add("firstRow", fmt.Sprintf("%v", *request.firstRow))
	}
	if request.firstColumn != nil {
		localVarQueryParams.Add("firstColumn", fmt.Sprintf("%v", *request.firstColumn))
	}
	if request.rowCount != nil {
		localVarQueryParams.Add("rowCount", fmt.Sprintf("%v", *request.rowCount))
	}
	if request.columnCount != nil {
		localVarQueryParams.Add("columnCount", fmt.Sprintf("%v", *request.columnCount))
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

func (request *GetWorksheetCellsRangeValueRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetWorksheetCellsRangeValueRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
