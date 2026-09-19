package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostAutofitWorksheetColumnsRequest struct {
	name      string
	sheetName string

	endColumn   *int
	folder      string
	onlyAuto    *bool
	startColumn *int
	storageName string

	extraQueryParameters map[string]string
}

func NewPostAutofitWorksheetColumnsRequest(name string, sheetName string, opts ...Option) *PostAutofitWorksheetColumnsRequest {
	req := &PostAutofitWorksheetColumnsRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["endColumn"].(*int); ok {
		req.endColumn = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["onlyAuto"].(*bool); ok {
		req.onlyAuto = val
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

func (request *PostAutofitWorksheetColumnsRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostAutofitWorksheetColumnsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostAutofitWorksheetColumnsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostAutofitWorksheetColumnsRequest) GetMethod() string {
	return "POST"
}

func (request *PostAutofitWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostAutofitWorksheetColumnsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autofitcolumns"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostAutofitWorksheetColumnsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.startColumn != nil {
		localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", *request.startColumn))
	}
	if request.endColumn != nil {
		localVarQueryParams.Add("endColumn", fmt.Sprintf("%v", *request.endColumn))
	}
	if request.onlyAuto != nil {
		localVarQueryParams.Add("onlyAuto", fmt.Sprintf("%v", *request.onlyAuto))
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

func (request *PostAutofitWorksheetColumnsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostAutofitWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
