package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type GenerateBatchReportRequest struct {
	data         string
	dataData     []byte
	dataName     string
	template     string
	templateData []byte
	templateName string

	AutoColumnsFit   *bool
	AutoRowsFit      *bool
	filename         string
	format           string
	mode             string
	outPath          string
	outStorageName   string
	password         string
	region           string
	reportObjectName string
	sheetName        string
	zipOutput        *bool

	extraQueryParameters map[string]string
}

func NewGenerateBatchReportRequest(data string, template string, opts ...Option) *GenerateBatchReportRequest {
	req := &GenerateBatchReportRequest{
		data:     data,
		template: template,
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
	if val, ok := cfg.Params["filename"].(string); ok {
		req.filename = val
	}
	if val, ok := cfg.Params["format"].(string); ok {
		req.format = val
	}
	if val, ok := cfg.Params["mode"].(string); ok {
		req.mode = val
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
	if val, ok := cfg.Params["reportObjectName"].(string); ok {
		req.reportObjectName = val
	}
	if val, ok := cfg.Params["sheetName"].(string); ok {
		req.sheetName = val
	}
	if val, ok := cfg.Params["zipOutput"].(*bool); ok {
		req.zipOutput = val
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

func (request *GenerateBatchReportRequest) Validate() error {
	if request.dataData == nil && request.data == "" {
		return fmt.Errorf("required request parameter %q is missing", "data")
	}

	if request.templateData == nil && request.template == "" {
		return fmt.Errorf("required request parameter %q is missing", "template")
	}

	return nil
}

func (request *GenerateBatchReportRequest) SetDataBytes(data []byte, name string) {
	if name == "" {
		name = "data"
	}
	request.dataData = data
	request.dataName = name
}

func (request *GenerateBatchReportRequest) SetTemplateBytes(data []byte, name string) {
	if name == "" {
		name = "template"
	}
	request.templateData = data
	request.templateName = name
}

func (request *GenerateBatchReportRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *GenerateBatchReportRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *GenerateBatchReportRequest) GetMethod() string {
	return "PUT"
}

func (request *GenerateBatchReportRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *GenerateBatchReportRequest) GetPath() string {
	localVarPath := "/v4.0/cells/report/generate/batch"
	return localVarPath
}

func (request *GenerateBatchReportRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.format != "" {
		localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
	}
	if request.mode != "" {
		localVarQueryParams.Add("mode", fmt.Sprintf("%v", request.mode))
	}
	if request.zipOutput != nil {
		localVarQueryParams.Add("zipOutput", fmt.Sprintf("%v", *request.zipOutput))
	}
	if request.reportObjectName != "" {
		localVarQueryParams.Add("reportObjectName", fmt.Sprintf("%v", request.reportObjectName))
	}
	if request.filename != "" {
		localVarQueryParams.Add("filename", fmt.Sprintf("%v", request.filename))
	}
	if request.sheetName != "" {
		localVarQueryParams.Add("sheetName", fmt.Sprintf("%v", request.sheetName))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.AutoRowsFit != nil {
		localVarQueryParams.Add("AutoRowsFit", fmt.Sprintf("%v", *request.AutoRowsFit))
	}
	if request.AutoColumnsFit != nil {
		localVarQueryParams.Add("AutoColumnsFit", fmt.Sprintf("%v", *request.AutoColumnsFit))
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

func (request *GenerateBatchReportRequest) GetJSONBody() interface{} {
	return nil
}

func (request *GenerateBatchReportRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.dataData != nil {
		localVarFormParams[request.dataName] = request.dataData
	} else if request.data != "" {
		localVarFormParams["@"+filepath.Base(request.data)] = request.data
	}
	if request.templateData != nil {
		localVarFormParams[request.templateName] = request.templateData
	} else if request.template != "" {
		localVarFormParams["@"+filepath.Base(request.template)] = request.template
	}
	return localVarFormParams
}
