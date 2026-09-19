package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type GetWorkbookRequest struct {
	name string

	checkExcelRestriction *bool
	folder                string
	FontsLocation         string
	format                string
	isAutoFit             *bool
	onePagePerSheet       *bool
	onlyAutofitTable      *bool
	onlySaveTable         *bool
	outPath               string
	outStorageName        string
	pageTallFitOnPerSheet *bool
	pageWideFitOnPerSheet *bool
	password              string
	region                string
	storageName           string

	extraQueryParameters map[string]string
}

func NewGetWorkbookRequest(name string, opts ...Option) *GetWorkbookRequest {
	req := &GetWorkbookRequest{
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
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["FontsLocation"].(string); ok {
		req.FontsLocation = val
	}
	if val, ok := cfg.Params["format"].(string); ok {
		req.format = val
	}
	if val, ok := cfg.Params["isAutoFit"].(*bool); ok {
		req.isAutoFit = val
	}
	if val, ok := cfg.Params["onePagePerSheet"].(*bool); ok {
		req.onePagePerSheet = val
	}
	if val, ok := cfg.Params["onlyAutofitTable"].(*bool); ok {
		req.onlyAutofitTable = val
	}
	if val, ok := cfg.Params["onlySaveTable"].(*bool); ok {
		req.onlySaveTable = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
	}
	if val, ok := cfg.Params["pageTallFitOnPerSheet"].(*bool); ok {
		req.pageTallFitOnPerSheet = val
	}
	if val, ok := cfg.Params["pageWideFitOnPerSheet"].(*bool); ok {
		req.pageWideFitOnPerSheet = val
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

func (request *GetWorkbookRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *GetWorkbookRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetWorkbookRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetWorkbookRequest) GetMethod() string {
	return "GET"
}

func (request *GetWorkbookRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetWorkbookRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *GetWorkbookRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.format != "" {
		localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.isAutoFit != nil {
		localVarQueryParams.Add("isAutoFit", fmt.Sprintf("%v", *request.isAutoFit))
	}
	if request.onlySaveTable != nil {
		localVarQueryParams.Add("onlySaveTable", fmt.Sprintf("%v", *request.onlySaveTable))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
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
	if request.pageWideFitOnPerSheet != nil {
		localVarQueryParams.Add("pageWideFitOnPerSheet", fmt.Sprintf("%v", *request.pageWideFitOnPerSheet))
	}
	if request.pageTallFitOnPerSheet != nil {
		localVarQueryParams.Add("pageTallFitOnPerSheet", fmt.Sprintf("%v", *request.pageTallFitOnPerSheet))
	}
	if request.onePagePerSheet != nil {
		localVarQueryParams.Add("onePagePerSheet", fmt.Sprintf("%v", *request.onePagePerSheet))
	}
	if request.onlyAutofitTable != nil {
		localVarQueryParams.Add("onlyAutofitTable", fmt.Sprintf("%v", *request.onlyAutofitTable))
	}
	if request.FontsLocation != "" {
		localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *GetWorkbookRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetWorkbookRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
