package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type ExportChartAsFormatRequest struct {
	chartIndex int
	format     string
	name       string
	worksheet  string

	folder         string
	fontsLocation  string
	outPath        string
	outStorageName string
	password       string
	region         string
	storageName    string

	extraQueryParameters map[string]string
}

func NewExportChartAsFormatRequest(chartIndex int, format string, name string, worksheet string, opts ...Option) *ExportChartAsFormatRequest {
	req := &ExportChartAsFormatRequest{
		chartIndex: chartIndex,
		format:     format,
		name:       name,
		worksheet:  worksheet,
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
	if val, ok := cfg.Params["fontsLocation"].(string); ok {
		req.fontsLocation = val
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

func (request *ExportChartAsFormatRequest) Validate() error {
	if request.format == "" {
		return fmt.Errorf("required request parameter %q is missing", "format")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *ExportChartAsFormatRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ExportChartAsFormatRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ExportChartAsFormatRequest) GetMethod() string {
	return "GET"
}

func (request *ExportChartAsFormatRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *ExportChartAsFormatRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/charts/{chartIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.chartIndex)), -1)
	return localVarPath
}

func (request *ExportChartAsFormatRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
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
	if request.fontsLocation != "" {
		localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
	}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *ExportChartAsFormatRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ExportChartAsFormatRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
