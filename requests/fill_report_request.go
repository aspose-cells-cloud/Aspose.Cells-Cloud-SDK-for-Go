package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type FillReportRequest struct {
	data     string
	dataData []byte
	dataName string

	columns        string
	filename       string
	format         string
	outPath        string
	outStorageName string
	password       string
	region         string

	extraQueryParameters map[string]string
}

func NewFillReportRequest(data string, opts ...Option) *FillReportRequest {
	req := &FillReportRequest{
		data: data,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["columns"].(string); ok {
		req.columns = val
	}
	if val, ok := cfg.Params["filename"].(string); ok {
		req.filename = val
	}
	if val, ok := cfg.Params["format"].(string); ok {
		req.format = val
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

func (request *FillReportRequest) Validate() error {
	if request.dataData == nil && request.data == "" {
		return fmt.Errorf("required request parameter %q is missing", "data")
	}

	return nil
}

func (request *FillReportRequest) SetDataBytes(data []byte, name string) {
	if name == "" {
		name = "data"
	}
	request.dataData = data
	request.dataName = name
}

func (request *FillReportRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *FillReportRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *FillReportRequest) GetMethod() string {
	return "PUT"
}

func (request *FillReportRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *FillReportRequest) GetPath() string {
	localVarPath := "/v4.0/cells/report/fill"
	return localVarPath
}

func (request *FillReportRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.format != "" {
		localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
	}
	if request.columns != "" {
		localVarQueryParams.Add("columns", fmt.Sprintf("%v", request.columns))
	}
	if request.filename != "" {
		localVarQueryParams.Add("filename", fmt.Sprintf("%v", request.filename))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
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

func (request *FillReportRequest) GetJSONBody() interface{} {
	return nil
}

func (request *FillReportRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.dataData != nil {
		localVarFormParams[request.dataName] = request.dataData
	} else if request.data != "" {
		localVarFormParams["@"+filepath.Base(request.data)] = request.data
	}
	return localVarFormParams
}
