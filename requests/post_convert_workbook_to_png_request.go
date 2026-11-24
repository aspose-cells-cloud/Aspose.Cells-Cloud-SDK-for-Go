package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PostConvertWorkbookToPNGRequest struct {
	File     string
	FileData []byte
	FileName string

	checkExcelRestriction *bool
	password              string
	region                string

	extraQueryParameters map[string]string
}

func NewPostConvertWorkbookToPNGRequest(File string, opts ...Option) *PostConvertWorkbookToPNGRequest {
	req := &PostConvertWorkbookToPNGRequest{
		File: File,
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

func (request *PostConvertWorkbookToPNGRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	return nil
}

func (request *PostConvertWorkbookToPNGRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostConvertWorkbookToPNGRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostConvertWorkbookToPNGRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostConvertWorkbookToPNGRequest) GetMethod() string {
	return "POST"
}

func (request *PostConvertWorkbookToPNGRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostConvertWorkbookToPNGRequest) GetPath() string {
	localVarPath := "/v3.0/cells/convert/png"
	return localVarPath
}

func (request *PostConvertWorkbookToPNGRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
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

func (request *PostConvertWorkbookToPNGRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostConvertWorkbookToPNGRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
