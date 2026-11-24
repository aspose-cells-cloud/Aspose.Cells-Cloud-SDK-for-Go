package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostSetCellRangeValueRequest struct {
	cellarea  string
	name      string
	sheetName string
	_type     string
	value     string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostSetCellRangeValueRequest(cellarea string, name string, sheetName string, _type string, value string, opts ...Option) *PostSetCellRangeValueRequest {
	req := &PostSetCellRangeValueRequest{
		cellarea:  cellarea,
		name:      name,
		sheetName: sheetName,
		_type:     _type,
		value:     value,
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

func (request *PostSetCellRangeValueRequest) Validate() error {
	if request.cellarea == "" {
		return fmt.Errorf("required request parameter %q is missing", "cellarea")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request._type == "" {
		return fmt.Errorf("required request parameter %q is missing", "type")
	}

	if request.value == "" {
		return fmt.Errorf("required request parameter %q is missing", "value")
	}

	return nil
}

func (request *PostSetCellRangeValueRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostSetCellRangeValueRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostSetCellRangeValueRequest) GetMethod() string {
	return "POST"
}

func (request *PostSetCellRangeValueRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostSetCellRangeValueRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/cells"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostSetCellRangeValueRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("cellarea", fmt.Sprintf("%v", request.cellarea))
	localVarQueryParams.Add("value", fmt.Sprintf("%v", request.value))
	localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
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

func (request *PostSetCellRangeValueRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostSetCellRangeValueRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
