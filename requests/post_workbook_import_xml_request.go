package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorkbookImportXMLRequest struct {
	importXMLRequest *models.ImportXMLRequest
	name             string

	checkExcelRestriction *bool
	folder                string
	outPath               string
	outStorageName        string
	password              string
	region                string
	storageName           string

	extraQueryParameters map[string]string
}

func NewPostWorkbookImportXMLRequest(importXMLRequest *models.ImportXMLRequest, name string, opts ...Option) *PostWorkbookImportXMLRequest {
	req := &PostWorkbookImportXMLRequest{
		importXMLRequest: importXMLRequest,
		name:             name,
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
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
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

func (request *PostWorkbookImportXMLRequest) Validate() error {
	if request.importXMLRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "importXMLRequest")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostWorkbookImportXMLRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookImportXMLRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookImportXMLRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookImportXMLRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookImportXMLRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/importxml"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbookImportXMLRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.checkExcelRestriction != nil {
		localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
	}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostWorkbookImportXMLRequest) GetJSONBody() interface{} {
	return &request.importXMLRequest
}

func (request *PostWorkbookImportXMLRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
