package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorkbookCreateRequest struct {
	name string

	checkExcelRestriction *bool
	dataFile              string
	folder                string
	isWriteOver           *bool
	storageName           string
	templateFile          string

	extraQueryParameters map[string]string
}

func NewPutWorkbookCreateRequest(name string, opts ...Option) *PutWorkbookCreateRequest {
	req := &PutWorkbookCreateRequest{
		name: name,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["checkExcelRestriction"].(*bool); ok {
		req.checkExcelRestriction = val
	}
	if val, ok := cfg.Params["dataFile"].(string); ok {
		req.dataFile = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["isWriteOver"].(*bool); ok {
		req.isWriteOver = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["templateFile"].(string); ok {
		req.templateFile = val
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

func (request *PutWorkbookCreateRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PutWorkbookCreateRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorkbookCreateRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorkbookCreateRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorkbookCreateRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorkbookCreateRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PutWorkbookCreateRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.templateFile != "" {
		localVarQueryParams.Add("templateFile", fmt.Sprintf("%v", request.templateFile))
	}
	if request.dataFile != "" {
		localVarQueryParams.Add("dataFile", fmt.Sprintf("%v", request.dataFile))
	}
	if request.isWriteOver != nil {
		localVarQueryParams.Add("isWriteOver", fmt.Sprintf("%v", *request.isWriteOver))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.checkExcelRestriction != nil {
		localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PutWorkbookCreateRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorkbookCreateRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
