package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostAutofitWorksheetRowsRequest struct {
	name      string
	sheetName string

	endRow      *int
	folder      string
	onlyAuto    *bool
	startRow    *int
	storageName string

	extraQueryParameters map[string]string
}

func NewPostAutofitWorksheetRowsRequest(name string, sheetName string, opts ...Option) *PostAutofitWorksheetRowsRequest {
	req := &PostAutofitWorksheetRowsRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["endRow"].(*int); ok {
		req.endRow = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["onlyAuto"].(*bool); ok {
		req.onlyAuto = val
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

func (request *PostAutofitWorksheetRowsRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostAutofitWorksheetRowsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostAutofitWorksheetRowsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostAutofitWorksheetRowsRequest) GetMethod() string {
	return "POST"
}

func (request *PostAutofitWorksheetRowsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostAutofitWorksheetRowsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autofitrows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostAutofitWorksheetRowsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.startRow != nil {
		localVarQueryParams.Add("startRow", fmt.Sprintf("%v", *request.startRow))
	}
	if request.endRow != nil {
		localVarQueryParams.Add("endRow", fmt.Sprintf("%v", *request.endRow))
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

func (request *PostAutofitWorksheetRowsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostAutofitWorksheetRowsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
