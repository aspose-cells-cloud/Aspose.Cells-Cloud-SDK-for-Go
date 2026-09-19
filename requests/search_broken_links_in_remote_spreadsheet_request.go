package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchBrokenLinksInRemoteSpreadsheetRequest struct {
	name string

	cellArea    string
	folder      string
	password    string
	region      string
	storageName string
	worksheet   string

	extraQueryParameters map[string]string
}

func NewSearchBrokenLinksInRemoteSpreadsheetRequest(name string, opts ...Option) *SearchBrokenLinksInRemoteSpreadsheetRequest {
	req := &SearchBrokenLinksInRemoteSpreadsheetRequest{
		name: name,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["cellArea"].(string); ok {
		req.cellArea = val
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
	if val, ok := cfg.Params["worksheet"].(string); ok {
		req.worksheet = val
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

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/search/broken-links"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.worksheet != "" {
		localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	}
	if request.cellArea != "" {
		localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
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

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SearchBrokenLinksInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
