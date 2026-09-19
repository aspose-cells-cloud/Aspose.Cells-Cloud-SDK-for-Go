package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostWorksheetListObjectInsertSlicerRequest struct {
	columnIndex     int
	destCellName    string
	listObjectIndex int
	name            string
	sheetName       string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorksheetListObjectInsertSlicerRequest(columnIndex int, destCellName string, listObjectIndex int, name string, sheetName string, opts ...Option) *PostWorksheetListObjectInsertSlicerRequest {
	req := &PostWorksheetListObjectInsertSlicerRequest{
		columnIndex:     columnIndex,
		destCellName:    destCellName,
		listObjectIndex: listObjectIndex,
		name:            name,
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

func (request *PostWorksheetListObjectInsertSlicerRequest) Validate() error {
	if request.destCellName == "" {
		return fmt.Errorf("required request parameter %q is missing", "destCellName")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostWorksheetListObjectInsertSlicerRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorksheetListObjectInsertSlicerRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/InsertSlicer"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.listObjectIndex)), -1)
	return localVarPath
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("columnIndex", fmt.Sprintf("%v", request.columnIndex))
	localVarQueryParams.Add("destCellName", fmt.Sprintf("%v", request.destCellName))
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

func (request *PostWorksheetListObjectInsertSlicerRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
