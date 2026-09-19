package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PutPivotTableFieldRequest struct {
	name                   string
	pivotFieldType         string
	pivotTableFieldRequest *models.PivotTableFieldRequest
	pivotTableIndex        int
	sheetName              string

	folder          string
	needReCalculate *bool
	storageName     string

	extraQueryParameters map[string]string
}

func NewPutPivotTableFieldRequest(name string, pivotFieldType string, pivotTableFieldRequest *models.PivotTableFieldRequest, pivotTableIndex int, sheetName string, opts ...Option) *PutPivotTableFieldRequest {
	req := &PutPivotTableFieldRequest{
		name:                   name,
		pivotFieldType:         pivotFieldType,
		pivotTableFieldRequest: pivotTableFieldRequest,
		pivotTableIndex:        pivotTableIndex,
		sheetName:              sheetName,
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

func (request *PutPivotTableFieldRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.pivotFieldType == "" {
		return fmt.Errorf("required request parameter %q is missing", "pivotFieldType")
	}

	if request.pivotTableFieldRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "pivotTableFieldRequest")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutPivotTableFieldRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutPivotTableFieldRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutPivotTableFieldRequest) GetMethod() string {
	return "PUT"
}

func (request *PutPivotTableFieldRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutPivotTableFieldRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.pivotTableIndex)), -1)
	return localVarPath
}

func (request *PutPivotTableFieldRequest) GetQueryParameters() url.Values {
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

func (request *PutPivotTableFieldRequest) GetJSONBody() interface{} {
	return &request.pivotTableFieldRequest
}

func (request *PutPivotTableFieldRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
