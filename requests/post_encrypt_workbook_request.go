package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostEncryptWorkbookRequest struct {
	encryption *models.WorkbookEncryptionRequest
	name       string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostEncryptWorkbookRequest(encryption *models.WorkbookEncryptionRequest, name string, opts ...Option) *PostEncryptWorkbookRequest {
	req := &PostEncryptWorkbookRequest{
		encryption: encryption,
		name:       name,
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

func (request *PostEncryptWorkbookRequest) Validate() error {
	if request.encryption == nil {
		return fmt.Errorf("required request parameter %q is missing", "encryption")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostEncryptWorkbookRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostEncryptWorkbookRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostEncryptWorkbookRequest) GetMethod() string {
	return "POST"
}

func (request *PostEncryptWorkbookRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostEncryptWorkbookRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/encryption"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostEncryptWorkbookRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
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

func (request *PostEncryptWorkbookRequest) GetJSONBody() interface{} {
	return &request.encryption
}

func (request *PostEncryptWorkbookRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
