package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PutConvertWorkbookRequest struct {
	File     string
	FileData []byte
	FileName string
	format   string

	AutoColumnsFit        *bool
	AutoRowsFit           *bool
	checkExcelRestriction *bool
	FontsLocation         string
	onePagePerSheet       *bool
	outPath               string
	pageIndex             *int
	pageTallFitOnPerSheet *bool
	pageWideFitOnPerSheet *bool
	password              string
	region                string
	sheetName             string
	storageName           string
	streamFormat          string

	extraQueryParameters map[string]string
}

func NewPutConvertWorkbookRequest(File string, format string, opts ...Option) *PutConvertWorkbookRequest {
	req := &PutConvertWorkbookRequest{
		File:   File,
		format: format,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["AutoColumnsFit"].(*bool); ok {
		req.AutoColumnsFit = val
	}
	if val, ok := cfg.Params["AutoRowsFit"].(*bool); ok {
		req.AutoRowsFit = val
	}
	if val, ok := cfg.Params["checkExcelRestriction"].(*bool); ok {
		req.checkExcelRestriction = val
	}
	if val, ok := cfg.Params["FontsLocation"].(string); ok {
		req.FontsLocation = val
	}
	if val, ok := cfg.Params["onePagePerSheet"].(*bool); ok {
		req.onePagePerSheet = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["pageIndex"].(*int); ok {
		req.pageIndex = val
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
	if val, ok := cfg.Params["sheetName"].(string); ok {
		req.sheetName = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["streamFormat"].(string); ok {
		req.streamFormat = val
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

func (request *PutConvertWorkbookRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	if request.format == "" {
		return fmt.Errorf("required request parameter %q is missing", "format")
	}

	return nil
}

func (request *PutConvertWorkbookRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PutConvertWorkbookRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutConvertWorkbookRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutConvertWorkbookRequest) GetMethod() string {
	return "PUT"
}

func (request *PutConvertWorkbookRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PutConvertWorkbookRequest) GetPath() string {
	localVarPath := "/v3.0/cells/convert"
	return localVarPath
}

func (request *PutConvertWorkbookRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.checkExcelRestriction != nil {
		localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
	}
	if request.streamFormat != "" {
		localVarQueryParams.Add("streamFormat", fmt.Sprintf("%v", request.streamFormat))
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
	if request.sheetName != "" {
		localVarQueryParams.Add("sheetName", fmt.Sprintf("%v", request.sheetName))
	}
	if request.pageIndex != nil {
		localVarQueryParams.Add("pageIndex", fmt.Sprintf("%v", *request.pageIndex))
	}
	if request.onePagePerSheet != nil {
		localVarQueryParams.Add("onePagePerSheet", fmt.Sprintf("%v", *request.onePagePerSheet))
	}
	if request.AutoRowsFit != nil {
		localVarQueryParams.Add("AutoRowsFit", fmt.Sprintf("%v", *request.AutoRowsFit))
	}
	if request.AutoColumnsFit != nil {
		localVarQueryParams.Add("AutoColumnsFit", fmt.Sprintf("%v", *request.AutoColumnsFit))
	}
	if request.FontsLocation != "" {
		localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PutConvertWorkbookRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutConvertWorkbookRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
