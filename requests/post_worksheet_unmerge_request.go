package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostWorksheetUnmergeRequest struct {
	name         string
	sheetName    string
	startColumn  int
	startRow     int
	totalColumns int
	totalRows    int

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostWorksheetUnmergeRequest(name string, sheetName string, startColumn int, startRow int, totalColumns int, totalRows int, opts ...Option) *PostWorksheetUnmergeRequest {
	req := &PostWorksheetUnmergeRequest{
		name:         name,
		sheetName:    sheetName,
		startColumn:  startColumn,
		startRow:     startRow,
		totalColumns: totalColumns,
		totalRows:    totalRows,
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

func (request *PostWorksheetUnmergeRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostWorksheetUnmergeRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorksheetUnmergeRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorksheetUnmergeRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorksheetUnmergeRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorksheetUnmergeRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/cells/unmerge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostWorksheetUnmergeRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("startRow", fmt.Sprintf("%v", request.startRow))
	localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", request.startColumn))
	localVarQueryParams.Add("totalRows", fmt.Sprintf("%v", request.totalRows))
	localVarQueryParams.Add("totalColumns", fmt.Sprintf("%v", request.totalColumns))
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

func (request *PostWorksheetUnmergeRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostWorksheetUnmergeRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
