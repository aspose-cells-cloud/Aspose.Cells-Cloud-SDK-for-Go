package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostProtectWorkbookRequest struct {
	name                   string
	protectWorkbookRequest *models.ProtectWorkbookRequest

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostProtectWorkbookRequest(name string, protectWorkbookRequest *models.ProtectWorkbookRequest, opts ...Option) *PostProtectWorkbookRequest {
	req := &PostProtectWorkbookRequest{
		name:                   name,
		protectWorkbookRequest: protectWorkbookRequest,
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

func (request *PostProtectWorkbookRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.protectWorkbookRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "protectWorkbookRequest")
	}

	return nil
}

func (request *PostProtectWorkbookRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostProtectWorkbookRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostProtectWorkbookRequest) GetMethod() string {
	return "POST"
}

func (request *PostProtectWorkbookRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostProtectWorkbookRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostProtectWorkbookRequest) GetQueryParameters() url.Values {
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

func (request *PostProtectWorkbookRequest) GetJSONBody() interface{} {
	return &request.protectWorkbookRequest
}

func (request *PostProtectWorkbookRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
