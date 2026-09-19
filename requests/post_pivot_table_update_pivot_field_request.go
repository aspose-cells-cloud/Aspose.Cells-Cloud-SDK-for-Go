package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostPivotTableUpdatePivotFieldRequest struct {
	name            string
	pivotField      *models.PivotField
	pivotFieldIndex int
	pivotFieldType  string
	pivotTableIndex int
	sheetName       string

	folder          string
	needReCalculate *bool
	storageName     string

	extraQueryParameters map[string]string
}

func NewPostPivotTableUpdatePivotFieldRequest(name string, pivotField *models.PivotField, pivotFieldIndex int, pivotFieldType string, pivotTableIndex int, sheetName string, opts ...Option) *PostPivotTableUpdatePivotFieldRequest {
	req := &PostPivotTableUpdatePivotFieldRequest{
		name:            name,
		pivotField:      pivotField,
		pivotFieldIndex: pivotFieldIndex,
		pivotFieldType:  pivotFieldType,
		pivotTableIndex: pivotTableIndex,
		sheetName:       sheetName,
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
	if val, ok := cfg.Params["needReCalculate"].(*bool); ok {
		req.needReCalculate = val
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

func (request *PostPivotTableUpdatePivotFieldRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.pivotField == nil {
		return fmt.Errorf("required request parameter %q is missing", "pivotField")
	}

	if request.pivotFieldType == "" {
		return fmt.Errorf("required request parameter %q is missing", "pivotFieldType")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostPivotTableUpdatePivotFieldRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostPivotTableUpdatePivotFieldRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostPivotTableUpdatePivotFieldRequest) GetMethod() string {
	return "POST"
}

func (request *PostPivotTableUpdatePivotFieldRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostPivotTableUpdatePivotFieldRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFields/{pivotFieldIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.pivotTableIndex)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotFieldIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.pivotFieldIndex)), -1)
	return localVarPath
}

func (request *PostPivotTableUpdatePivotFieldRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("pivotFieldType", fmt.Sprintf("%v", request.pivotFieldType))
	if request.needReCalculate != nil {
		localVarQueryParams.Add("needReCalculate", fmt.Sprintf("%v", *request.needReCalculate))
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

func (request *PostPivotTableUpdatePivotFieldRequest) GetJSONBody() interface{} {
	return &request.pivotField
}

func (request *PostPivotTableUpdatePivotFieldRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
