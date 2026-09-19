package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostPivotTableStyleRequest struct {
	name            string
	pivotTableIndex int
	sheetName       string
	style           *models.Style

	folder          string
	needReCalculate *bool
	storageName     string

	extraQueryParameters map[string]string
}

func NewPostPivotTableStyleRequest(name string, pivotTableIndex int, sheetName string, style *models.Style, opts ...Option) *PostPivotTableStyleRequest {
	req := &PostPivotTableStyleRequest{
		name:            name,
		pivotTableIndex: pivotTableIndex,
		sheetName:       sheetName,
		style:           style,
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

func (request *PostPivotTableStyleRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request.style == nil {
		return fmt.Errorf("required request parameter %q is missing", "style")
	}

	return nil
}

func (request *PostPivotTableStyleRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostPivotTableStyleRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostPivotTableStyleRequest) GetMethod() string {
	return "POST"
}

func (request *PostPivotTableStyleRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostPivotTableStyleRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/FormatAll"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.pivotTableIndex)), -1)
	return localVarPath
}

func (request *PostPivotTableStyleRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
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

func (request *PostPivotTableStyleRequest) GetJSONBody() interface{} {
	return &request.style
}

func (request *PostPivotTableStyleRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
