package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostDigitalSignatureRequest struct {
	digitalsignaturefile string
	name                 string
	password             string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostDigitalSignatureRequest(digitalsignaturefile string, name string, password string, opts ...Option) *PostDigitalSignatureRequest {
	req := &PostDigitalSignatureRequest{
		digitalsignaturefile: digitalsignaturefile,
		name:                 name,
		password:             password,
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

func (request *PostDigitalSignatureRequest) Validate() error {
	if request.digitalsignaturefile == "" {
		return fmt.Errorf("required request parameter %q is missing", "digitalsignaturefile")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.password == "" {
		return fmt.Errorf("required request parameter %q is missing", "password")
	}

	return nil
}

func (request *PostDigitalSignatureRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostDigitalSignatureRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostDigitalSignatureRequest) GetMethod() string {
	return "POST"
}

func (request *PostDigitalSignatureRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostDigitalSignatureRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/digitalsignature"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostDigitalSignatureRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("digitalsignaturefile", fmt.Sprintf("%v", request.digitalsignaturefile))
	localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
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

func (request *PostDigitalSignatureRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostDigitalSignatureRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
