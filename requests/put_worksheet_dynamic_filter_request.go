package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetDynamicFilterRequest struct {
	dynamicFilterType string
	fieldIndex        int
	name              string
	_range            string
	sheetName         string

	folder      string
	matchBlanks *bool
	refresh     *bool
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetDynamicFilterRequest(dynamicFilterType string, fieldIndex int, name string, _range string, sheetName string, opts ...Option) *PutWorksheetDynamicFilterRequest {
	req := &PutWorksheetDynamicFilterRequest{
		dynamicFilterType: dynamicFilterType,
		fieldIndex:        fieldIndex,
		name:              name,
		_range:            _range,
		sheetName:         sheetName,
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
	if val, ok := cfg.Params["matchBlanks"].(*bool); ok {
		req.matchBlanks = val
	}
	if val, ok := cfg.Params["refresh"].(*bool); ok {
		req.refresh = val
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

func (request *PutWorksheetDynamicFilterRequest) Validate() error {
	if request.dynamicFilterType == "" {
		return fmt.Errorf("required request parameter %q is missing", "dynamicFilterType")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetDynamicFilterRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetDynamicFilterRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetDynamicFilterRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetDynamicFilterRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetDynamicFilterRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autoFilter/dynamicFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetDynamicFilterRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
	localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
	localVarQueryParams.Add("dynamicFilterType", fmt.Sprintf("%v", request.dynamicFilterType))
	if request.matchBlanks != nil {
		localVarQueryParams.Add("matchBlanks", fmt.Sprintf("%v", *request.matchBlanks))
	}
	if request.refresh != nil {
		localVarQueryParams.Add("refresh", fmt.Sprintf("%v", *request.refresh))
	}
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

func (request *PutWorksheetDynamicFilterRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetDynamicFilterRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
