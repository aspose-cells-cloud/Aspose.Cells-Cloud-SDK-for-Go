package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PutDocumentProtectFromChangesRequest struct {
	name     string
	password *models.PasswordRequest

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutDocumentProtectFromChangesRequest(name string, password *models.PasswordRequest, opts ...Option) *PutDocumentProtectFromChangesRequest {
	req := &PutDocumentProtectFromChangesRequest{
		name:     name,
		password: password,
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

func (request *PutDocumentProtectFromChangesRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.password == nil {
		return fmt.Errorf("required request parameter %q is missing", "password")
	}

	return nil
}

func (request *PutDocumentProtectFromChangesRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutDocumentProtectFromChangesRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutDocumentProtectFromChangesRequest) GetMethod() string {
	return "PUT"
}

func (request *PutDocumentProtectFromChangesRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutDocumentProtectFromChangesRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/writeProtection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PutDocumentProtectFromChangesRequest) GetQueryParameters() url.Values {
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

func (request *PutDocumentProtectFromChangesRequest) GetJSONBody() interface{} {
	return &request.password
}

func (request *PutDocumentProtectFromChangesRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
