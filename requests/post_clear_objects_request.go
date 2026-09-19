package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PostClearObjectsRequest struct {
	File       string
	FileData   []byte
	FileName   string
	objecttype string

	checkExcelRestriction *bool
	outFormat             string
	password              string
	region                string
	sheetname             string

	extraQueryParameters map[string]string
}

func NewPostClearObjectsRequest(File string, objecttype string, opts ...Option) *PostClearObjectsRequest {
	req := &PostClearObjectsRequest{
		File:       File,
		objecttype: objecttype,
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
	if val, ok := cfg.Params["outFormat"].(string); ok {
		req.outFormat = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["sheetname"].(string); ok {
		req.sheetname = val
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

func (request *PostClearObjectsRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	if request.objecttype == "" {
		return fmt.Errorf("required request parameter %q is missing", "objecttype")
	}

	return nil
}

func (request *PostClearObjectsRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostClearObjectsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostClearObjectsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostClearObjectsRequest) GetMethod() string {
	return "POST"
}

func (request *PostClearObjectsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostClearObjectsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/clearobjects"
	return localVarPath
}

func (request *PostClearObjectsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("objecttype", fmt.Sprintf("%v", request.objecttype))
	if request.sheetname != "" {
		localVarQueryParams.Add("sheetname", fmt.Sprintf("%v", request.sheetname))
	}
	if request.outFormat != "" {
		localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
	}
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

func (request *PostClearObjectsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostClearObjectsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
