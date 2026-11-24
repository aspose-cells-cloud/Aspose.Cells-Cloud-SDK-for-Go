package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type DeletePivotTableFieldRequest struct {
	name                   string
	pivotFieldType         string
	pivotTableFieldRequest *models.PivotTableFieldRequest
	pivotTableIndex        int
	sheetName              string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewDeletePivotTableFieldRequest(name string, pivotFieldType string, pivotTableFieldRequest *models.PivotTableFieldRequest, pivotTableIndex int, sheetName string, opts ...Option) *DeletePivotTableFieldRequest {
	req := &DeletePivotTableFieldRequest{
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

func (request *DeletePivotTableFieldRequest) Validate() error {
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

func (request *DeletePivotTableFieldRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *DeletePivotTableFieldRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *DeletePivotTableFieldRequest) GetMethod() string {
	return "DELETE"
}

func (request *DeletePivotTableFieldRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *DeletePivotTableFieldRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.pivotTableIndex)), -1)
	return localVarPath
}

func (request *DeletePivotTableFieldRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("pivotFieldType", fmt.Sprintf("%v", request.pivotFieldType))
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

func (request *DeletePivotTableFieldRequest) GetJSONBody() interface{} {
	return &request.pivotTableFieldRequest
}

func (request *DeletePivotTableFieldRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
