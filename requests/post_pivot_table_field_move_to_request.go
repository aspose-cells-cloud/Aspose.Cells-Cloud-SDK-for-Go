package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostPivotTableFieldMoveToRequest struct {
	fieldIndex      int
	from            string
	name            string
	pivotTableIndex int
	sheetName       string
	to              string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostPivotTableFieldMoveToRequest(fieldIndex int, from string, name string, pivotTableIndex int, sheetName string, to string, opts ...Option) *PostPivotTableFieldMoveToRequest {
	req := &PostPivotTableFieldMoveToRequest{
		fieldIndex:      fieldIndex,
		from:            from,
		name:            name,
		pivotTableIndex: pivotTableIndex,
		sheetName:       sheetName,
		to:              to,
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

func (request *PostPivotTableFieldMoveToRequest) Validate() error {
	if request.from == "" {
		return fmt.Errorf("required request parameter %q is missing", "from")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request.to == "" {
		return fmt.Errorf("required request parameter %q is missing", "to")
	}

	return nil
}

func (request *PostPivotTableFieldMoveToRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostPivotTableFieldMoveToRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostPivotTableFieldMoveToRequest) GetMethod() string {
	return "POST"
}

func (request *PostPivotTableFieldMoveToRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostPivotTableFieldMoveToRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Move"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.pivotTableIndex)), -1)
	return localVarPath
}

func (request *PostPivotTableFieldMoveToRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
	localVarQueryParams.Add("from", fmt.Sprintf("%v", request.from))
	localVarQueryParams.Add("to", fmt.Sprintf("%v", request.to))
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

func (request *PostPivotTableFieldMoveToRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostPivotTableFieldMoveToRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
