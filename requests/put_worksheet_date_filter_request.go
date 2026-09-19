package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetDateFilterRequest struct {
	dateTimeGroupingType string
	fieldIndex           int
	name                 string
	_range               string
	sheetName            string

	day         *int
	folder      string
	hour        *int
	matchBlanks *bool
	minute      *int
	month       *int
	refresh     *bool
	second      *int
	storageName string
	year        *int

	extraQueryParameters map[string]string
}

func NewPutWorksheetDateFilterRequest(dateTimeGroupingType string, fieldIndex int, name string, _range string, sheetName string, opts ...Option) *PutWorksheetDateFilterRequest {
	req := &PutWorksheetDateFilterRequest{
		dateTimeGroupingType: dateTimeGroupingType,
		fieldIndex:           fieldIndex,
		name:                 name,
		_range:               _range,
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
	if val, ok := cfg.Params["matchBlanks"].(*bool); ok {
		req.matchBlanks = val
	}
	if val, ok := cfg.Params["minute"].(*int); ok {
		req.minute = val
	}
	if val, ok := cfg.Params["month"].(*int); ok {
		req.month = val
	}
	if val, ok := cfg.Params["refresh"].(*bool); ok {
		req.refresh = val
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

func (request *PutWorksheetDateFilterRequest) Validate() error {
	if request.dateTimeGroupingType == "" {
		return fmt.Errorf("required request parameter %q is missing", "dateTimeGroupingType")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetDateFilterRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetDateFilterRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetDateFilterRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetDateFilterRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetDateFilterRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/autoFilter/dateFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetDateFilterRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
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
	if request.matchBlanks != nil {
		localVarQueryParams.Add("matchBlanks", fmt.Sprintf("%v", *request.matchBlanks))
	}
	if request.refresh != nil {
		localVarQueryParams.Add("refresh", fmt.Sprintf("%v", *request.refresh))
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

func (request *PutWorksheetDateFilterRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetDateFilterRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
