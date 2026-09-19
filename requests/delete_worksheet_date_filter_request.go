package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type DeleteWorksheetDateFilterRequest struct {
	dateTimeGroupingType string
	fieldIndex           int
	name                 string
	sheetName            string

	day         *int
	folder      string
	hour        *int
	minute      *int
	month       *int
	second      *int
	storageName string
	year        *int

	extraQueryParameters map[string]string
}

func NewDeleteWorksheetDateFilterRequest(dateTimeGroupingType string, fieldIndex int, name string, sheetName string, opts ...Option) *DeleteWorksheetDateFilterRequest {
	req := &DeleteWorksheetDateFilterRequest{
		dateTimeGroupingType: dateTimeGroupingType,
		fieldIndex:           fieldIndex,
		name:                 name,
		sheetName:            sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["day"].(*int); ok {
		req.day = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["hour"].(*int); ok {
		req.hour = val
	}
	if val, ok := cfg.Params["minute"].(*int); ok {
		req.minute = val
	}
	if val, ok := cfg.Params["month"].(*int); ok {
		req.month = val
	}
	if val, ok := cfg.Params["second"].(*int); ok {
		req.second = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["year"].(*int); ok {
		req.year = val
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

func (request *DeleteWorksheetDateFilterRequest) Validate() error {
	if request.dateTimeGroupingType == "" {
		return fmt.Errorf("required request parameter %q is missing", "dateTimeGroupingType")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *DeleteWorksheetDateFilterRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *DeleteWorksheetDateFilterRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *DeleteWorksheetDateFilterRequest) GetMethod() string {
	return "DELETE"
}

func (request *DeleteWorksheetDateFilterRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *DeleteWorksheetDateFilterRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autoFilter/dateFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *DeleteWorksheetDateFilterRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
	localVarQueryParams.Add("dateTimeGroupingType", fmt.Sprintf("%v", request.dateTimeGroupingType))
	if request.year != nil {
		localVarQueryParams.Add("year", fmt.Sprintf("%v", *request.year))
	}
	if request.month != nil {
		localVarQueryParams.Add("month", fmt.Sprintf("%v", *request.month))
	}
	if request.day != nil {
		localVarQueryParams.Add("day", fmt.Sprintf("%v", *request.day))
	}
	if request.hour != nil {
		localVarQueryParams.Add("hour", fmt.Sprintf("%v", *request.hour))
	}
	if request.minute != nil {
		localVarQueryParams.Add("minute", fmt.Sprintf("%v", *request.minute))
	}
	if request.second != nil {
		localVarQueryParams.Add("second", fmt.Sprintf("%v", *request.second))
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

func (request *DeleteWorksheetDateFilterRequest) GetJSONBody() interface{} {
	return nil
}

func (request *DeleteWorksheetDateFilterRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
