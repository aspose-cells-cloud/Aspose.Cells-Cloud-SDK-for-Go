package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutInsertNewWorksheetRequest struct {
	index     int
	name      string
	sheetName string
	sheettype string

	folder       string
	newsheetname string
	storageName  string

	extraQueryParameters map[string]string
}

func NewPutInsertNewWorksheetRequest(index int, name string, sheetName string, sheettype string, opts ...Option) *PutInsertNewWorksheetRequest {
	req := &PutInsertNewWorksheetRequest{
		index:     index,
		name:      name,
		sheetName: sheetName,
		sheettype: sheettype,
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
	if val, ok := cfg.Params["newsheetname"].(string); ok {
		req.newsheetname = val
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

func (request *PutInsertNewWorksheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request.sheettype == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheettype")
	}

	return nil
}

func (request *PutInsertNewWorksheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutInsertNewWorksheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutInsertNewWorksheetRequest) GetMethod() string {
	return "PUT"
}

func (request *PutInsertNewWorksheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutInsertNewWorksheetRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/insert"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PutInsertNewWorksheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("sheetName", fmt.Sprintf("%v", request.sheetName))
	localVarQueryParams.Add("index", fmt.Sprintf("%v", request.index))
	localVarQueryParams.Add("sheettype", fmt.Sprintf("%v", request.sheettype))
	if request.newsheetname != "" {
		localVarQueryParams.Add("newsheetname", fmt.Sprintf("%v", request.newsheetname))
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

func (request *PutInsertNewWorksheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutInsertNewWorksheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
