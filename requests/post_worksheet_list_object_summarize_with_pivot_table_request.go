package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorksheetListObjectSummarizeWithPivotTableRequest struct {
	createPivotTableRequest *models.CreatePivotTableRequest
	destsheetName           string
	listObjectIndex         int
	name                    string
	sheetName               string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorksheetListObjectSummarizeWithPivotTableRequest(createPivotTableRequest *models.CreatePivotTableRequest, destsheetName string, listObjectIndex int, name string, sheetName string, opts ...Option) *PostWorksheetListObjectSummarizeWithPivotTableRequest {
	req := &PostWorksheetListObjectSummarizeWithPivotTableRequest{
		createPivotTableRequest: createPivotTableRequest,
		destsheetName:           destsheetName,
		listObjectIndex:         listObjectIndex,
		name:                    name,
		sheetName:               sheetName,
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

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) Validate() error {
	if request.createPivotTableRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "createPivotTableRequest")
	}

	if request.destsheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "destsheetName")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/SummarizeWithPivotTable"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.listObjectIndex)), -1)
	return localVarPath
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("destsheetName", fmt.Sprintf("%v", request.destsheetName))
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

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetJSONBody() interface{} {
	return &request.createPivotTableRequest
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
