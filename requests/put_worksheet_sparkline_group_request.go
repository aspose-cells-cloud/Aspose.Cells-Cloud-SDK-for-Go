package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetSparklineGroupRequest struct {
	dataRange     string
	isVertical    bool
	locationRange string
	name          string
	sheetName     string
	_type         string

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorksheetSparklineGroupRequest(dataRange string, isVertical bool, locationRange string, name string, sheetName string, _type string, opts ...Option) *PutWorksheetSparklineGroupRequest {
	req := &PutWorksheetSparklineGroupRequest{
		dataRange:     dataRange,
		isVertical:    isVertical,
		locationRange: locationRange,
		name:          name,
		sheetName:     sheetName,
		_type:         _type,
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

func (request *PutWorksheetSparklineGroupRequest) Validate() error {
	if request.dataRange == "" {
		return fmt.Errorf("required request parameter %q is missing", "dataRange")
	}

	if request.locationRange == "" {
		return fmt.Errorf("required request parameter %q is missing", "locationRange")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request._type == "" {
		return fmt.Errorf("required request parameter %q is missing", "type")
	}

	return nil
}

func (request *PutWorksheetSparklineGroupRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetSparklineGroupRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetSparklineGroupRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetSparklineGroupRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetSparklineGroupRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/sparklineGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetSparklineGroupRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
	localVarQueryParams.Add("dataRange", fmt.Sprintf("%v", request.dataRange))
	localVarQueryParams.Add("isVertical", fmt.Sprintf("%v", request.isVertical))
	localVarQueryParams.Add("locationRange", fmt.Sprintf("%v", request.locationRange))
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

func (request *PutWorksheetSparklineGroupRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetSparklineGroupRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
