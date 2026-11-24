package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PostSetWorksheetColumnWidthRequest struct {
	columnIndex int
	name        string
	sheetName   string
	width       float64

	count       *int
	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPostSetWorksheetColumnWidthRequest(columnIndex int, name string, sheetName string, width float64, opts ...Option) *PostSetWorksheetColumnWidthRequest {
	req := &PostSetWorksheetColumnWidthRequest{
		columnIndex: columnIndex,
		name:        name,
		sheetName:   sheetName,
		width:       width,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["count"].(*int); ok {
		req.count = val
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

func (request *PostSetWorksheetColumnWidthRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PostSetWorksheetColumnWidthRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostSetWorksheetColumnWidthRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostSetWorksheetColumnWidthRequest) GetMethod() string {
	return "POST"
}

func (request *PostSetWorksheetColumnWidthRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostSetWorksheetColumnWidthRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", url.PathEscape(fmt.Sprintf("%v", request.columnIndex)), -1)
	return localVarPath
}

func (request *PostSetWorksheetColumnWidthRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("width", fmt.Sprintf("%v", request.width))
	if request.count != nil {
		localVarQueryParams.Add("count", fmt.Sprintf("%v", *request.count))
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

func (request *PostSetWorksheetColumnWidthRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostSetWorksheetColumnWidthRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
