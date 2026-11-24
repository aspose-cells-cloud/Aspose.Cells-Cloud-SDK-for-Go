package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetCustomFilterRequest struct {
	criteria1     string
	fieldIndex    int
	name          string
	operatorType1 string
	_range        string
	sheetName     string

	criteria2     string
	folder        string
	isAnd         *bool
	matchBlanks   *bool
	operatorType2 string
	refresh       *bool
	storageName   string

	extraQueryParameters map[string]string
}

func NewPutWorksheetCustomFilterRequest(criteria1 string, fieldIndex int, name string, operatorType1 string, _range string, sheetName string, opts ...Option) *PutWorksheetCustomFilterRequest {
	req := &PutWorksheetCustomFilterRequest{
		criteria1:     criteria1,
		fieldIndex:    fieldIndex,
		name:          name,
		operatorType1: operatorType1,
		_range:        _range,
		sheetName:     sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["criteria2"].(string); ok {
		req.criteria2 = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["isAnd"].(*bool); ok {
		req.isAnd = val
	}
	if val, ok := cfg.Params["matchBlanks"].(*bool); ok {
		req.matchBlanks = val
	}
	if val, ok := cfg.Params["operatorType2"].(string); ok {
		req.operatorType2 = val
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

func (request *PutWorksheetCustomFilterRequest) Validate() error {
	if request.criteria1 == "" {
		return fmt.Errorf("required request parameter %q is missing", "criteria1")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.operatorType1 == "" {
		return fmt.Errorf("required request parameter %q is missing", "operatorType1")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetCustomFilterRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetCustomFilterRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetCustomFilterRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetCustomFilterRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetCustomFilterRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autoFilter/custom"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetCustomFilterRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
	localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
	localVarQueryParams.Add("operatorType1", fmt.Sprintf("%v", request.operatorType1))
	localVarQueryParams.Add("criteria1", fmt.Sprintf("%v", request.criteria1))
	if request.isAnd != nil {
		localVarQueryParams.Add("isAnd", fmt.Sprintf("%v", *request.isAnd))
	}
	if request.operatorType2 != "" {
		localVarQueryParams.Add("operatorType2", fmt.Sprintf("%v", request.operatorType2))
	}
	if request.criteria2 != "" {
		localVarQueryParams.Add("criteria2", fmt.Sprintf("%v", request.criteria2))
	}
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

func (request *PutWorksheetCustomFilterRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetCustomFilterRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
