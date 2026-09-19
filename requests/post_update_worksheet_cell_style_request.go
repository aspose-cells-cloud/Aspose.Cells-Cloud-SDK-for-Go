package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostUpdateWorksheetCellStyleRequest struct {
	cellName  string
	name      string
	sheetName string
	style     *models.Style

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostUpdateWorksheetCellStyleRequest(cellName string, name string, sheetName string, style *models.Style, opts ...Option) *PostUpdateWorksheetCellStyleRequest {
	req := &PostUpdateWorksheetCellStyleRequest{
		cellName:  cellName,
		name:      name,
		sheetName: sheetName,
		style:     style,
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

func (request *PostUpdateWorksheetCellStyleRequest) Validate() error {
	if request.cellName == "" {
		return fmt.Errorf("required request parameter %q is missing", "cellName")
	}

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

func (request *PostUpdateWorksheetCellStyleRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostUpdateWorksheetCellStyleRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostUpdateWorksheetCellStyleRequest) GetMethod() string {
	return "POST"
}

func (request *PostUpdateWorksheetCellStyleRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostUpdateWorksheetCellStyleRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/cells/{cellName}/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", url.PathEscape(fmt.Sprintf("%v", request.cellName)), -1)
	return localVarPath
}

func (request *PostUpdateWorksheetCellStyleRequest) GetQueryParameters() url.Values {
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

func (request *PostUpdateWorksheetCellStyleRequest) GetJSONBody() interface{} {
	return &request.style
}

func (request *PostUpdateWorksheetCellStyleRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
