package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutInsertWorksheetColumnsRequest struct {
	columnIndex int
	columns     int
	name        string
	sheetName   string

	folder          string
	storageName     string
	updateReference *bool

	extraQueryParameters map[string]string
}

func NewPutInsertWorksheetColumnsRequest(columnIndex int, columns int, name string, sheetName string, opts ...Option) *PutInsertWorksheetColumnsRequest {
	req := &PutInsertWorksheetColumnsRequest{
		columnIndex: columnIndex,
		columns:     columns,
		name:        name,
		sheetName:   sheetName,
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
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["updateReference"].(*bool); ok {
		req.updateReference = val
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

func (request *PutInsertWorksheetColumnsRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutInsertWorksheetColumnsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutInsertWorksheetColumnsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutInsertWorksheetColumnsRequest) GetMethod() string {
	return "PUT"
}

func (request *PutInsertWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutInsertWorksheetColumnsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.columnIndex)), -1)
	return localVarPath
}

func (request *PutInsertWorksheetColumnsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("columns", fmt.Sprintf("%v", request.columns))
	if request.updateReference != nil {
		localVarQueryParams.Add("updateReference", fmt.Sprintf("%v", *request.updateReference))
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

func (request *PutInsertWorksheetColumnsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutInsertWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
