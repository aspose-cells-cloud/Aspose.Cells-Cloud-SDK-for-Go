package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type ConvertTextInRemoteSpreadsheetRequest struct {
	convertTextType string
	name            string
	_range          string
	worksheet       string

	folder           string
	password         string
	region           string
	sourceCharacters string
	storageName      string
	targetCharacters string

	extraQueryParameters map[string]string
}

func NewConvertTextInRemoteSpreadsheetRequest(convertTextType string, name string, _range string, worksheet string, opts ...Option) *ConvertTextInRemoteSpreadsheetRequest {
	req := &ConvertTextInRemoteSpreadsheetRequest{
		convertTextType: convertTextType,
		name:            name,
		_range:          _range,
		worksheet:       worksheet,
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
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["sourceCharacters"].(string); ok {
		req.sourceCharacters = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["targetCharacters"].(string); ok {
		req.targetCharacters = val
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

func (request *ConvertTextInRemoteSpreadsheetRequest) Validate() error {
	if request.convertTextType == "" {
		return fmt.Errorf("required request parameter %q is missing", "convertTextType")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *ConvertTextInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *ConvertTextInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/range/{range}/content/convert/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", url.PathEscape(fmt.Sprintf("%v", request._range)), -1)
	return localVarPath
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("convertTextType", fmt.Sprintf("%v", request.convertTextType))
	if request.sourceCharacters != "" {
		localVarQueryParams.Add("sourceCharacters", fmt.Sprintf("%v", request.sourceCharacters))
	}
	if request.targetCharacters != "" {
		localVarQueryParams.Add("targetCharacters", fmt.Sprintf("%v", request.targetCharacters))
	}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
