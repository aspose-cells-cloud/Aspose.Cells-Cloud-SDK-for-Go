package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutAddNewWorksheetRequest struct {
	name      string
	sheetName string

	folder      string
	position    *int
	sheettype   string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutAddNewWorksheetRequest(name string, sheetName string, opts ...Option) *PutAddNewWorksheetRequest {
	req := &PutAddNewWorksheetRequest{
		name:      name,
		sheetName: sheetName,
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
	if val, ok := cfg.Params["position"].(*int); ok {
		req.position = val
	}
	if val, ok := cfg.Params["sheettype"].(string); ok {
		req.sheettype = val
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

func (request *PutAddNewWorksheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutAddNewWorksheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutAddNewWorksheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutAddNewWorksheetRequest) GetMethod() string {
	return "PUT"
}

func (request *PutAddNewWorksheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutAddNewWorksheetRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutAddNewWorksheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.position != nil {
		localVarQueryParams.Add("position", fmt.Sprintf("%v", *request.position))
	}
	if request.sheettype != "" {
		localVarQueryParams.Add("sheettype", fmt.Sprintf("%v", request.sheettype))
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

func (request *PutAddNewWorksheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutAddNewWorksheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
