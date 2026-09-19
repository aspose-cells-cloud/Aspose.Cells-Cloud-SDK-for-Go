package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type GetWorksheetWithFormatRequest struct {
	name      string
	sheetName string

	area                 string
	folder               string
	format               string
	horizontalResolution *int
	onePagePerSheet      *bool
	pageIndex            *int
	printHeadings        *bool
	storageName          string
	verticalResolution   *int

	extraQueryParameters map[string]string
}

func NewGetWorksheetWithFormatRequest(name string, sheetName string, opts ...Option) *GetWorksheetWithFormatRequest {
	req := &GetWorksheetWithFormatRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["area"].(string); ok {
		req.area = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["format"].(string); ok {
		req.format = val
	}
	if val, ok := cfg.Params["horizontalResolution"].(*int); ok {
		req.horizontalResolution = val
	}
	if val, ok := cfg.Params["onePagePerSheet"].(*bool); ok {
		req.onePagePerSheet = val
	}
	if val, ok := cfg.Params["pageIndex"].(*int); ok {
		req.pageIndex = val
	}
	if val, ok := cfg.Params["printHeadings"].(*bool); ok {
		req.printHeadings = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["verticalResolution"].(*int); ok {
		req.verticalResolution = val
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

func (request *GetWorksheetWithFormatRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *GetWorksheetWithFormatRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GetWorksheetWithFormatRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GetWorksheetWithFormatRequest) GetMethod() string {
	return "GET"
}

func (request *GetWorksheetWithFormatRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *GetWorksheetWithFormatRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *GetWorksheetWithFormatRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.format != "" {
		localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
	}
	if request.verticalResolution != nil {
		localVarQueryParams.Add("verticalResolution", fmt.Sprintf("%v", *request.verticalResolution))
	}
	if request.horizontalResolution != nil {
		localVarQueryParams.Add("horizontalResolution", fmt.Sprintf("%v", *request.horizontalResolution))
	}
	if request.area != "" {
		localVarQueryParams.Add("area", fmt.Sprintf("%v", request.area))
	}
	if request.pageIndex != nil {
		localVarQueryParams.Add("pageIndex", fmt.Sprintf("%v", *request.pageIndex))
	}
	if request.onePagePerSheet != nil {
		localVarQueryParams.Add("onePagePerSheet", fmt.Sprintf("%v", *request.onePagePerSheet))
	}
	if request.printHeadings != nil {
		localVarQueryParams.Add("printHeadings", fmt.Sprintf("%v", *request.printHeadings))
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

func (request *GetWorksheetWithFormatRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GetWorksheetWithFormatRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
