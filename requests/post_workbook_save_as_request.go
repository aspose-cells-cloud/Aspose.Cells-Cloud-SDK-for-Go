package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorkbookSaveAsRequest struct {
	name        string
	newfilename string

	checkExcelRestriction *bool
	folder                string
	FontsLocation         string
	isAutoFitColumns      *bool
	isAutoFitRows         *bool
	onePagePerSheet       *bool
	outStorageName        string
	pageTallFitOnPerSheet *bool
	pageWideFitOnPerSheet *bool
	region                string
	saveOptions           models.SaveOptionsLike
	storageName           string

	extraQueryParameters map[string]string
}

func NewPostWorkbookSaveAsRequest(name string, newfilename string, opts ...Option) *PostWorkbookSaveAsRequest {
	req := &PostWorkbookSaveAsRequest{
		name:        name,
		newfilename: newfilename,
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
	if val, ok := cfg.Params["isAutoFitColumns"].(*bool); ok {
		req.isAutoFitColumns = val
	}
	if val, ok := cfg.Params["isAutoFitRows"].(*bool); ok {
		req.isAutoFitRows = val
	}
	if val, ok := cfg.Params["onePagePerSheet"].(*bool); ok {
		req.onePagePerSheet = val
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
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["saveOptions"].(models.SaveOptionsLike); ok {
		req.saveOptions = val
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

func (request *PostWorkbookSaveAsRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.newfilename == "" {
		return fmt.Errorf("required request parameter %q is missing", "newfilename")
	}

	return nil
}

func (request *PostWorkbookSaveAsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookSaveAsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookSaveAsRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookSaveAsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookSaveAsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/SaveAs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbookSaveAsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("newfilename", fmt.Sprintf("%v", request.newfilename))
	if request.isAutoFitRows != nil {
		localVarQueryParams.Add("isAutoFitRows", fmt.Sprintf("%v", *request.isAutoFitRows))
	}
	if request.isAutoFitColumns != nil {
		localVarQueryParams.Add("isAutoFitColumns", fmt.Sprintf("%v", *request.isAutoFitColumns))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
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
	if request.FontsLocation != "" {
		localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostWorkbookSaveAsRequest) GetJSONBody() interface{} {
	return &request.saveOptions
}

func (request *PostWorkbookSaveAsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
