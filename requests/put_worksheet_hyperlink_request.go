package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetHyperlinkRequest struct {
	address      string
	firstColumn  int
	firstRow     int
	name         string
	sheetName    string
	totalColumns int
	totalRows    int

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetHyperlinkRequest(address string, firstColumn int, firstRow int, name string, sheetName string, totalColumns int, totalRows int, opts ...Option) *PutWorksheetHyperlinkRequest {
	req := &PutWorksheetHyperlinkRequest{
		address:      address,
		firstColumn:  firstColumn,
		firstRow:     firstRow,
		name:         name,
		sheetName:    sheetName,
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

func (request *PutWorksheetHyperlinkRequest) Validate() error {
	if request.address == "" {
		return fmt.Errorf("required request parameter %q is missing", "address")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetHyperlinkRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetHyperlinkRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetHyperlinkRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetHyperlinkRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetHyperlinkRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/hyperlinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetHyperlinkRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("firstRow", fmt.Sprintf("%v", request.firstRow))
	localVarQueryParams.Add("firstColumn", fmt.Sprintf("%v", request.firstColumn))
	localVarQueryParams.Add("totalRows", fmt.Sprintf("%v", request.totalRows))
	localVarQueryParams.Add("totalColumns", fmt.Sprintf("%v", request.totalColumns))
	localVarQueryParams.Add("address", fmt.Sprintf("%v", request.address))
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

func (request *PutWorksheetHyperlinkRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetHyperlinkRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
