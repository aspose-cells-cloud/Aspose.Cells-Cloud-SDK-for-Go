package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostAutofitWorkbookRowsRequest struct {
	name string

	endRow      *int
	firstColumn *int
	folder      string
	lastColumn  *int
	onlyAuto    *bool
	startRow    *int
	storageName string

	extraQueryParameters map[string]string
}

func NewPostAutofitWorkbookRowsRequest(name string, opts ...Option) *PostAutofitWorkbookRowsRequest {
	req := &PostAutofitWorkbookRowsRequest{
		name: name,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["endRow"].(*int); ok {
		req.endRow = val
	}
	if val, ok := cfg.Params["firstColumn"].(*int); ok {
		req.firstColumn = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["lastColumn"].(*int); ok {
		req.lastColumn = val
	}
	if val, ok := cfg.Params["onlyAuto"].(*bool); ok {
		req.onlyAuto = val
	}
	if val, ok := cfg.Params["startRow"].(*int); ok {
		req.startRow = val
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

func (request *PostAutofitWorkbookRowsRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostAutofitWorkbookRowsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostAutofitWorkbookRowsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostAutofitWorkbookRowsRequest) GetMethod() string {
	return "POST"
}

func (request *PostAutofitWorkbookRowsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostAutofitWorkbookRowsRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/autofitrows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostAutofitWorkbookRowsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.startRow != nil {
		localVarQueryParams.Add("startRow", fmt.Sprintf("%v", *request.startRow))
	}
	if request.endRow != nil {
		localVarQueryParams.Add("endRow", fmt.Sprintf("%v", *request.endRow))
	}
	if request.onlyAuto != nil {
		localVarQueryParams.Add("onlyAuto", fmt.Sprintf("%v", *request.onlyAuto))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.firstColumn != nil {
		localVarQueryParams.Add("firstColumn", fmt.Sprintf("%v", *request.firstColumn))
	}
	if request.lastColumn != nil {
		localVarQueryParams.Add("lastColumn", fmt.Sprintf("%v", *request.lastColumn))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostAutofitWorkbookRowsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostAutofitWorkbookRowsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
