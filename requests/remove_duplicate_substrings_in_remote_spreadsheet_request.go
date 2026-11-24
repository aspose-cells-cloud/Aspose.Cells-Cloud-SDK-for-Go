package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type RemoveDuplicateSubstringsInRemoteSpreadsheetRequest struct {
	delimiters string
	name       string
	_range     string
	worksheet  string

	caseSensitive                   *bool
	folder                          string
	password                        string
	region                          string
	storageName                     string
	treatConsecutiveDelimitersAsOne *bool

	extraQueryParameters map[string]string
}

func NewRemoveDuplicateSubstringsInRemoteSpreadsheetRequest(delimiters string, name string, _range string, worksheet string, opts ...Option) *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest {
	req := &RemoveDuplicateSubstringsInRemoteSpreadsheetRequest{
		delimiters: delimiters,
		name:       name,
		_range:     _range,
		worksheet:  worksheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["caseSensitive"].(*bool); ok {
		req.caseSensitive = val
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
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["treatConsecutiveDelimitersAsOne"].(*bool); ok {
		req.treatConsecutiveDelimitersAsOne = val
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

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) Validate() error {
	if request.delimiters == "" {
		return fmt.Errorf("required request parameter %q is missing", "delimiters")
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

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/range/{range}/content/remove/duplicate-substrings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", url.PathEscape(fmt.Sprintf("%v", request._range)), -1)
	return localVarPath
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("delimiters", fmt.Sprintf("%v", request.delimiters))
	if request.treatConsecutiveDelimitersAsOne != nil {
		localVarQueryParams.Add("treatConsecutiveDelimitersAsOne", fmt.Sprintf("%v", *request.treatConsecutiveDelimitersAsOne))
	}
	if request.caseSensitive != nil {
		localVarQueryParams.Add("caseSensitive", fmt.Sprintf("%v", *request.caseSensitive))
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

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
