package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostGroupWorksheetColumnsRequest struct {
	firstIndex int
	lastIndex  int
	name       string
	sheetName  string

	folder      string
	hide        *bool
	storageName string

	extraQueryParameters map[string]string
}

func NewPostGroupWorksheetColumnsRequest(firstIndex int, lastIndex int, name string, sheetName string, opts ...Option) *PostGroupWorksheetColumnsRequest {
	req := &PostGroupWorksheetColumnsRequest{
		firstIndex: firstIndex,
		lastIndex:  lastIndex,
		name:       name,
		sheetName:  sheetName,
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
	if val, ok := cfg.Params["hide"].(*bool); ok {
		req.hide = val
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

func (request *PostGroupWorksheetColumnsRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostGroupWorksheetColumnsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostGroupWorksheetColumnsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostGroupWorksheetColumnsRequest) GetMethod() string {
	return "POST"
}

func (request *PostGroupWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostGroupWorksheetColumnsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/cells/columns/group"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostGroupWorksheetColumnsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("firstIndex", fmt.Sprintf("%v", request.firstIndex))
	localVarQueryParams.Add("lastIndex", fmt.Sprintf("%v", request.lastIndex))
	if request.hide != nil {
		localVarQueryParams.Add("hide", fmt.Sprintf("%v", *request.hide))
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

func (request *PostGroupWorksheetColumnsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostGroupWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
