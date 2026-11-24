package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PostSplitRequest struct {
	File      string
	FileData  []byte
	FileName  string
	outFormat string

	checkExcelRestriction *bool
	from                  *int
	password              string
	region                string
	to                    *int

	extraQueryParameters map[string]string
}

func NewPostSplitRequest(File string, outFormat string, opts ...Option) *PostSplitRequest {
	req := &PostSplitRequest{
		File:      File,
		outFormat: outFormat,
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
	if val, ok := cfg.Params["from"].(*int); ok {
		req.from = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["to"].(*int); ok {
		req.to = val
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

func (request *PostSplitRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	if request.outFormat == "" {
		return fmt.Errorf("required request parameter %q is missing", "outFormat")
	}

	return nil
}

func (request *PostSplitRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostSplitRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostSplitRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostSplitRequest) GetMethod() string {
	return "POST"
}

func (request *PostSplitRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostSplitRequest) GetPath() string {
	localVarPath := "/v3.0/cells/split"
	return localVarPath
}

func (request *PostSplitRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.from != nil {
		localVarQueryParams.Add("from", fmt.Sprintf("%v", *request.from))
	}
	if request.to != nil {
		localVarQueryParams.Add("to", fmt.Sprintf("%v", *request.to))
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

func (request *PostSplitRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostSplitRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
