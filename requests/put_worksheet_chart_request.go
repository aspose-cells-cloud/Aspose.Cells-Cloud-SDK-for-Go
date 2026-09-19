package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetChartRequest struct {
	chartType string
	name      string
	sheetName string

	area                string
	categoryData        string
	dataLabels          *bool
	dataLabelsPosition  string
	folder              string
	isAutoGetSerialName *bool
	isVertical          *bool
	lowerRightColumn    *int
	lowerRightRow       *int
	pivotTableName      string
	pivotTableSheet     string
	storageName         string
	title               string
	upperLeftColumn     *int
	upperLeftRow        *int

	extraQueryParameters map[string]string
}

func NewPutWorksheetChartRequest(chartType string, name string, sheetName string, opts ...Option) *PutWorksheetChartRequest {
	req := &PutWorksheetChartRequest{
		chartType: chartType,
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["area"].(string); ok {
		req.area = val
	}
	if val, ok := cfg.Params["categoryData"].(string); ok {
		req.categoryData = val
	}
	if val, ok := cfg.Params["dataLabels"].(*bool); ok {
		req.dataLabels = val
	}
	if val, ok := cfg.Params["dataLabelsPosition"].(string); ok {
		req.dataLabelsPosition = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["isAutoGetSerialName"].(*bool); ok {
		req.isAutoGetSerialName = val
	}
	if val, ok := cfg.Params["isVertical"].(*bool); ok {
		req.isVertical = val
	}
	if val, ok := cfg.Params["lowerRightColumn"].(*int); ok {
		req.lowerRightColumn = val
	}
	if val, ok := cfg.Params["lowerRightRow"].(*int); ok {
		req.lowerRightRow = val
	}
	if val, ok := cfg.Params["pivotTableName"].(string); ok {
		req.pivotTableName = val
	}
	if val, ok := cfg.Params["pivotTableSheet"].(string); ok {
		req.pivotTableSheet = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["title"].(string); ok {
		req.title = val
	}
	if val, ok := cfg.Params["upperLeftColumn"].(*int); ok {
		req.upperLeftColumn = val
	}
	if val, ok := cfg.Params["upperLeftRow"].(*int); ok {
		req.upperLeftRow = val
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

func (request *PutWorksheetChartRequest) Validate() error {
	if request.chartType == "" {
		return fmt.Errorf("required request parameter %q is missing", "chartType")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetChartRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetChartRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetChartRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetChartRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetChartRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/charts"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetChartRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("chartType", fmt.Sprintf("%v", request.chartType))
	if request.upperLeftRow != nil {
		localVarQueryParams.Add("upperLeftRow", fmt.Sprintf("%v", *request.upperLeftRow))
	}
	if request.upperLeftColumn != nil {
		localVarQueryParams.Add("upperLeftColumn", fmt.Sprintf("%v", *request.upperLeftColumn))
	}
	if request.lowerRightRow != nil {
		localVarQueryParams.Add("lowerRightRow", fmt.Sprintf("%v", *request.lowerRightRow))
	}
	if request.lowerRightColumn != nil {
		localVarQueryParams.Add("lowerRightColumn", fmt.Sprintf("%v", *request.lowerRightColumn))
	}
	if request.area != "" {
		localVarQueryParams.Add("area", fmt.Sprintf("%v", request.area))
	}
	if request.isVertical != nil {
		localVarQueryParams.Add("isVertical", fmt.Sprintf("%v", *request.isVertical))
	}
	if request.categoryData != "" {
		localVarQueryParams.Add("categoryData", fmt.Sprintf("%v", request.categoryData))
	}
	if request.isAutoGetSerialName != nil {
		localVarQueryParams.Add("isAutoGetSerialName", fmt.Sprintf("%v", *request.isAutoGetSerialName))
	}
	if request.title != "" {
		localVarQueryParams.Add("title", fmt.Sprintf("%v", request.title))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.dataLabels != nil {
		localVarQueryParams.Add("dataLabels", fmt.Sprintf("%v", *request.dataLabels))
	}
	if request.dataLabelsPosition != "" {
		localVarQueryParams.Add("dataLabelsPosition", fmt.Sprintf("%v", request.dataLabelsPosition))
	}
	if request.pivotTableSheet != "" {
		localVarQueryParams.Add("pivotTableSheet", fmt.Sprintf("%v", request.pivotTableSheet))
	}
	if request.pivotTableName != "" {
		localVarQueryParams.Add("pivotTableName", fmt.Sprintf("%v", request.pivotTableName))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PutWorksheetChartRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetChartRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
