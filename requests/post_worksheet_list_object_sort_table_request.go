package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorksheetListObjectSortTableRequest struct {
	dataSorter      *models.DataSorter
	listObjectIndex int
	name            string
	sheetName       string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorksheetListObjectSortTableRequest(dataSorter *models.DataSorter, listObjectIndex int, name string, sheetName string, opts ...Option) *PostWorksheetListObjectSortTableRequest {
	req := &PostWorksheetListObjectSortTableRequest{
		dataSorter:      dataSorter,
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

func (request *PostWorksheetListObjectSortTableRequest) Validate() error {
	if request.dataSorter == nil {
		return fmt.Errorf("required request parameter %q is missing", "dataSorter")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostWorksheetListObjectSortTableRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorksheetListObjectSortTableRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorksheetListObjectSortTableRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorksheetListObjectSortTableRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorksheetListObjectSortTableRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/sort"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.listObjectIndex)), -1)
	return localVarPath
}

func (request *PostWorksheetListObjectSortTableRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
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

func (request *PostWorksheetListObjectSortTableRequest) GetJSONBody() interface{} {
	return &request.dataSorter
}

func (request *PostWorksheetListObjectSortTableRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
