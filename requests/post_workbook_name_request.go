package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorkbookNameRequest struct {
	name     string
	nameName string
	newName  *models.Name

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorkbookNameRequest(name string, nameName string, newName *models.Name, opts ...Option) *PostWorkbookNameRequest {
	req := &PostWorkbookNameRequest{
		name:     name,
		nameName: nameName,
		newName:  newName,
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

func (request *PostWorkbookNameRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.nameName == "" {
		return fmt.Errorf("required request parameter %q is missing", "nameName")
	}

	if request.newName == nil {
		return fmt.Errorf("required request parameter %q is missing", "newName")
	}

	return nil
}

func (request *PostWorkbookNameRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookNameRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookNameRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookNameRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookNameRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/names/{nameName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", url.PathEscape(fmt.Sprintf("%v", request.nameName)), -1)
	return localVarPath
}

func (request *PostWorkbookNameRequest) GetQueryParameters() url.Values {
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

func (request *PostWorkbookNameRequest) GetJSONBody() interface{} {
	return &request.newName
}

func (request *PostWorkbookNameRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
