package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostAutofitWorksheetRowRequest struct {
	name      string
	rowIndex  int
	sheetName string

	firstColumn *int
	folder      string
	lastColumn  *int
	rowCount    *int
	storageName string

	extraQueryParameters map[string]string
}

func NewPostAutofitWorksheetRowRequest(name string, rowIndex int, sheetName string, opts ...Option) *PostAutofitWorksheetRowRequest {
	req := &PostAutofitWorksheetRowRequest{
		name:      name,
		rowIndex:  rowIndex,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["firstColumn"].(*int); ok {
		req.firstColumn = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["lastColumn"].(*int); ok {
		req.lastColumn = val
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

func (request *PostAutofitWorksheetRowRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostAutofitWorksheetRowRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostAutofitWorksheetRowRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostAutofitWorksheetRowRequest) GetMethod() string {
	return "POST"
}

func (request *PostAutofitWorksheetRowRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostAutofitWorksheetRowRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autofitrow"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostAutofitWorksheetRowRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("rowIndex", fmt.Sprintf("%v", request.rowIndex))
	if request.firstColumn != nil {
		localVarQueryParams.Add("firstColumn", fmt.Sprintf("%v", *request.firstColumn))
	}
	if request.lastColumn != nil {
		localVarQueryParams.Add("lastColumn", fmt.Sprintf("%v", *request.lastColumn))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.rowCount != nil {
		localVarQueryParams.Add("rowCount", fmt.Sprintf("%v", *request.rowCount))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostAutofitWorksheetRowRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostAutofitWorksheetRowRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
