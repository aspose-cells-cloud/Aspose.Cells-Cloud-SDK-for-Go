package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetFormatConditionAreaRequest struct {
	cellArea  string
	index     int
	name      string
	sheetName string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetFormatConditionAreaRequest(cellArea string, index int, name string, sheetName string, opts ...Option) *PutWorksheetFormatConditionAreaRequest {
	req := &PutWorksheetFormatConditionAreaRequest{
		cellArea:  cellArea,
		index:     index,
		name:      name,
		sheetName: sheetName,
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

func (request *PutWorksheetFormatConditionAreaRequest) Validate() error {
	if request.cellArea == "" {
		return fmt.Errorf("required request parameter %q is missing", "cellArea")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetFormatConditionAreaRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetFormatConditionAreaRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetFormatConditionAreaRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetFormatConditionAreaRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetFormatConditionAreaRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/area"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", url.PathEscape(fmt.Sprintf("%v", request.index)), -1)
	return localVarPath
}

func (request *PutWorksheetFormatConditionAreaRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
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

func (request *PutWorksheetFormatConditionAreaRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetFormatConditionAreaRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
