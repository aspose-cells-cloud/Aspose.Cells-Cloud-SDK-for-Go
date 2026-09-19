package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PostUnlockRequest struct {
	File     string
	FileData []byte
	FileName string
	password string

	extraQueryParameters map[string]string
}

func NewPostUnlockRequest(File string, password string, opts ...Option) *PostUnlockRequest {
	req := &PostUnlockRequest{
		File:     File,
		password: password,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
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

func (request *PostUnlockRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	if request.password == "" {
		return fmt.Errorf("required request parameter %q is missing", "password")
	}

	return nil
}

func (request *PostUnlockRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostUnlockRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostUnlockRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostUnlockRequest) GetMethod() string {
	return "POST"
}

func (request *PostUnlockRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostUnlockRequest) GetPath() string {
	localVarPath := "/v3.0/cells/unlock"
	return localVarPath
}

func (request *PostUnlockRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostUnlockRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostUnlockRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
