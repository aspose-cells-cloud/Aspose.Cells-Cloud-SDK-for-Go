package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostWorkbookGetSmartMarkerResultRequest struct {
	name string

	folder         string
	outPath        string
	outStorageName string
	storageName    string
	xmlFile        string

	extraQueryParameters map[string]string
}

func NewPostWorkbookGetSmartMarkerResultRequest(name string, opts ...Option) *PostWorkbookGetSmartMarkerResultRequest {
	req := &PostWorkbookGetSmartMarkerResultRequest{
		name: name,
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
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["xmlFile"].(string); ok {
		req.xmlFile = val
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

func (request *PostWorkbookGetSmartMarkerResultRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostWorkbookGetSmartMarkerResultRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookGetSmartMarkerResultRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookGetSmartMarkerResultRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookGetSmartMarkerResultRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookGetSmartMarkerResultRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/smartmarker"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbookGetSmartMarkerResultRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.xmlFile != "" {
		localVarQueryParams.Add("xmlFile", fmt.Sprintf("%v", request.xmlFile))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostWorkbookGetSmartMarkerResultRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostWorkbookGetSmartMarkerResultRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
