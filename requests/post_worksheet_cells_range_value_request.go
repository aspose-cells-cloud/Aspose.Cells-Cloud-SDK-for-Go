package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorksheetCellsRangeValueRequest struct {
	name      string
	_range    *models.Range
	sheetName string
	Value     string

	folder      string
	isConverted *bool
	setStyle    *bool
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorksheetCellsRangeValueRequest(name string, _range *models.Range, sheetName string, Value string, opts ...Option) *PostWorksheetCellsRangeValueRequest {
	req := &PostWorksheetCellsRangeValueRequest{
		name:      name,
		_range:    _range,
		sheetName: sheetName,
		Value:     Value,
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
	if val, ok := cfg.Params["isConverted"].(*bool); ok {
		req.isConverted = val
	}
	if val, ok := cfg.Params["setStyle"].(*bool); ok {
		req.setStyle = val
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

func (request *PostWorksheetCellsRangeValueRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request._range == nil {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request.Value == "" {
		return fmt.Errorf("required request parameter %q is missing", "Value")
	}

	return nil
}

func (request *PostWorksheetCellsRangeValueRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorksheetCellsRangeValueRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorksheetCellsRangeValueRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorksheetCellsRangeValueRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeValueRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/ranges/value"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostWorksheetCellsRangeValueRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("Value", fmt.Sprintf("%v", request.Value))
	if request.isConverted != nil {
		localVarQueryParams.Add("isConverted", fmt.Sprintf("%v", *request.isConverted))
	}
	if request.setStyle != nil {
		localVarQueryParams.Add("setStyle", fmt.Sprintf("%v", *request.setStyle))
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

func (request *PostWorksheetCellsRangeValueRequest) GetJSONBody() interface{} {
	return &request._range
}

func (request *PostWorksheetCellsRangeValueRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
