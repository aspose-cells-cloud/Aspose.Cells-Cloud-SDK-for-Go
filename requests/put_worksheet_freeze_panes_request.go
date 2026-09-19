package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetFreezePanesRequest struct {
	column         int
	freezedColumns int
	freezedRows    int
	name           string
	row            int
	sheetName      string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetFreezePanesRequest(column int, freezedColumns int, freezedRows int, name string, row int, sheetName string, opts ...Option) *PutWorksheetFreezePanesRequest {
	req := &PutWorksheetFreezePanesRequest{
		column:         column,
		freezedColumns: freezedColumns,
		freezedRows:    freezedRows,
		name:           name,
		row:            row,
		sheetName:      sheetName,
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

func (request *PutWorksheetFreezePanesRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetFreezePanesRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetFreezePanesRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetFreezePanesRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetFreezePanesRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetFreezePanesRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/freezepanes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetFreezePanesRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("row", fmt.Sprintf("%v", request.row))
	localVarQueryParams.Add("column", fmt.Sprintf("%v", request.column))
	localVarQueryParams.Add("freezedRows", fmt.Sprintf("%v", request.freezedRows))
	localVarQueryParams.Add("freezedColumns", fmt.Sprintf("%v", request.freezedColumns))
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

func (request *PutWorksheetFreezePanesRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetFreezePanesRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
