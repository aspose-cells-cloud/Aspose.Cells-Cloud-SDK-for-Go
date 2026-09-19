package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostWorkbookTextReplaceRequest struct {
	name     string
	newValue string
	oldValue string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorkbookTextReplaceRequest(name string, newValue string, oldValue string, opts ...Option) *PostWorkbookTextReplaceRequest {
	req := &PostWorkbookTextReplaceRequest{
		name:     name,
		newValue: newValue,
		oldValue: oldValue,
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

func (request *PostWorkbookTextReplaceRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.newValue == "" {
		return fmt.Errorf("required request parameter %q is missing", "newValue")
	}

	if request.oldValue == "" {
		return fmt.Errorf("required request parameter %q is missing", "oldValue")
	}

	return nil
}

func (request *PostWorkbookTextReplaceRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookTextReplaceRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookTextReplaceRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookTextReplaceRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookTextReplaceRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/replaceText"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbookTextReplaceRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("oldValue", fmt.Sprintf("%v", request.oldValue))
	localVarQueryParams.Add("newValue", fmt.Sprintf("%v", request.newValue))
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

func (request *PostWorkbookTextReplaceRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostWorkbookTextReplaceRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
