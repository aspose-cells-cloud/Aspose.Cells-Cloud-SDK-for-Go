package requests

import (
	"fmt"
	"net/url"
	"path/filepath"

	"asposecellscloud/models"
)

type PostProtectRequest struct {
	File                   string
	FileData               []byte
	FileName               string
	protectWorkbookRequest *models.ProtectWorkbookRequest

	password string

	extraQueryParameters map[string]string
}

func NewPostProtectRequest(File string, protectWorkbookRequest *models.ProtectWorkbookRequest, opts ...Option) *PostProtectRequest {
	req := &PostProtectRequest{
		File:                   File,
		protectWorkbookRequest: protectWorkbookRequest,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
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

func (request *PostProtectRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	if request.protectWorkbookRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "protectWorkbookRequest")
	}

	return nil
}

func (request *PostProtectRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostProtectRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostProtectRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostProtectRequest) GetMethod() string {
	return "POST"
}

func (request *PostProtectRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostProtectRequest) GetPath() string {
	localVarPath := "/v3.0/cells/protect"
	return localVarPath
}

func (request *PostProtectRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostProtectRequest) GetJSONBody() interface{} {
	return &request.protectWorkbookRequest
}

func (request *PostProtectRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
