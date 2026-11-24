package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetListObjectRequest struct {
	name      string
	sheetName string

	displayName string
	endColumn   *int
	endRow      *int
	folder      string
	hasHeaders  *bool
	showTotals  *bool
	startColumn *int
	startRow    *int
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetListObjectRequest(name string, sheetName string, opts ...Option) *PutWorksheetListObjectRequest {
	req := &PutWorksheetListObjectRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["displayName"].(string); ok {
		req.displayName = val
	}
	if val, ok := cfg.Params["endColumn"].(*int); ok {
		req.endColumn = val
	}
	if val, ok := cfg.Params["endRow"].(*int); ok {
		req.endRow = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["hasHeaders"].(*bool); ok {
		req.hasHeaders = val
	}
	if val, ok := cfg.Params["showTotals"].(*bool); ok {
		req.showTotals = val
	}
	if val, ok := cfg.Params["startColumn"].(*int); ok {
		req.startColumn = val
	}
	if val, ok := cfg.Params["startRow"].(*int); ok {
		req.startRow = val
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

func (request *PutWorksheetListObjectRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetListObjectRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetListObjectRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetListObjectRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetListObjectRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetListObjectRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/listobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetListObjectRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.startRow != nil {
		localVarQueryParams.Add("startRow", fmt.Sprintf("%v", *request.startRow))
	}
	if request.startColumn != nil {
		localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", *request.startColumn))
	}
	if request.endRow != nil {
		localVarQueryParams.Add("endRow", fmt.Sprintf("%v", *request.endRow))
	}
	if request.endColumn != nil {
		localVarQueryParams.Add("endColumn", fmt.Sprintf("%v", *request.endColumn))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.hasHeaders != nil {
		localVarQueryParams.Add("hasHeaders", fmt.Sprintf("%v", *request.hasHeaders))
	}
	if request.displayName != "" {
		localVarQueryParams.Add("displayName", fmt.Sprintf("%v", request.displayName))
	}
	if request.showTotals != nil {
		localVarQueryParams.Add("showTotals", fmt.Sprintf("%v", *request.showTotals))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PutWorksheetListObjectRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetListObjectRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
