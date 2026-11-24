package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutHorizontalPageBreakRequest struct {
	name      string
	sheetName string

	cellname    string
	column      *int
	endColumn   *int
	folder      string
	row         *int
	startColumn *int
	storageName string

	extraQueryParameters map[string]string
}

func NewPutHorizontalPageBreakRequest(name string, sheetName string, opts ...Option) *PutHorizontalPageBreakRequest {
	req := &PutHorizontalPageBreakRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["cellname"].(string); ok {
		req.cellname = val
	}
	if val, ok := cfg.Params["column"].(*int); ok {
		req.column = val
	}
	if val, ok := cfg.Params["endColumn"].(*int); ok {
		req.endColumn = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["row"].(*int); ok {
		req.row = val
	}
	if val, ok := cfg.Params["startColumn"].(*int); ok {
		req.startColumn = val
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

func (request *PutHorizontalPageBreakRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutHorizontalPageBreakRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutHorizontalPageBreakRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutHorizontalPageBreakRequest) GetMethod() string {
	return "PUT"
}

func (request *PutHorizontalPageBreakRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutHorizontalPageBreakRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutHorizontalPageBreakRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.cellname != "" {
		localVarQueryParams.Add("cellname", fmt.Sprintf("%v", request.cellname))
	}
	if request.row != nil {
		localVarQueryParams.Add("row", fmt.Sprintf("%v", *request.row))
	}
	if request.column != nil {
		localVarQueryParams.Add("column", fmt.Sprintf("%v", *request.column))
	}
	if request.startColumn != nil {
		localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", *request.startColumn))
	}
	if request.endColumn != nil {
		localVarQueryParams.Add("endColumn", fmt.Sprintf("%v", *request.endColumn))
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

func (request *PutHorizontalPageBreakRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutHorizontalPageBreakRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
