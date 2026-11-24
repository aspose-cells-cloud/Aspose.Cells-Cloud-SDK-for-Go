package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetFormatConditionRequest struct {
	cellArea     string
	formula1     string
	formula2     string
	index        int
	name         string
	operatorType string
	sheetName    string
	_type        string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetFormatConditionRequest(cellArea string, formula1 string, formula2 string, index int, name string, operatorType string, sheetName string, _type string, opts ...Option) *PutWorksheetFormatConditionRequest {
	req := &PutWorksheetFormatConditionRequest{
		cellArea:     cellArea,
		formula1:     formula1,
		formula2:     formula2,
		index:        index,
		name:         name,
		operatorType: operatorType,
		sheetName:    sheetName,
		_type:        _type,
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

func (request *PutWorksheetFormatConditionRequest) Validate() error {
	if request.cellArea == "" {
		return fmt.Errorf("required request parameter %q is missing", "cellArea")
	}

	if request.formula1 == "" {
		return fmt.Errorf("required request parameter %q is missing", "formula1")
	}

	if request.formula2 == "" {
		return fmt.Errorf("required request parameter %q is missing", "formula2")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.operatorType == "" {
		return fmt.Errorf("required request parameter %q is missing", "operatorType")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request._type == "" {
		return fmt.Errorf("required request parameter %q is missing", "type")
	}

	return nil
}

func (request *PutWorksheetFormatConditionRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetFormatConditionRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetFormatConditionRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetFormatConditionRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetFormatConditionRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", url.PathEscape(fmt.Sprintf("%v", request.index)), -1)
	return localVarPath
}

func (request *PutWorksheetFormatConditionRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
	localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
	localVarQueryParams.Add("operatorType", fmt.Sprintf("%v", request.operatorType))
	localVarQueryParams.Add("formula1", fmt.Sprintf("%v", request.formula1))
	localVarQueryParams.Add("formula2", fmt.Sprintf("%v", request.formula2))
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

func (request *PutWorksheetFormatConditionRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetFormatConditionRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
