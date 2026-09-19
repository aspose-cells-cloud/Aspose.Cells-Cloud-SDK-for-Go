package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PostReplaceRequest struct {
	File     string
	FileData []byte
	FileName string
	newtext  string
	text     string

	checkExcelRestriction *bool
	password              string
	sheetname             string

	extraQueryParameters map[string]string
}

func NewPostReplaceRequest(File string, newtext string, text string, opts ...Option) *PostReplaceRequest {
	req := &PostReplaceRequest{
		File:    File,
		newtext: newtext,
		text:    text,
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

func (request *PostReplaceRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	if request.newtext == "" {
		return fmt.Errorf("required request parameter %q is missing", "newtext")
	}

	if request.text == "" {
		return fmt.Errorf("required request parameter %q is missing", "text")
	}

	return nil
}

func (request *PostReplaceRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostReplaceRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostReplaceRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostReplaceRequest) GetMethod() string {
	return "POST"
}

func (request *PostReplaceRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostReplaceRequest) GetPath() string {
	localVarPath := "/v3.0/cells/replace"
	return localVarPath
}

func (request *PostReplaceRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
	localVarQueryParams.Add("newtext", fmt.Sprintf("%v", request.newtext))
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.sheetname != "" {
		localVarQueryParams.Add("sheetname", fmt.Sprintf("%v", request.sheetname))
	}
	if request.checkExcelRestriction != nil {
		localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostReplaceRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostReplaceRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
