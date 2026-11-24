package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type MergeRemoteSpreadsheetRequest struct {
	mergedSpreadsheet string
	name              string

	folder          string
	fontsLocation   string
	mergeInOneSheet *bool
	outFormat       string
	outPath         string
	outStorageName  string
	password        string
	region          string
	storageName     string

	extraQueryParameters map[string]string
}

func NewMergeRemoteSpreadsheetRequest(mergedSpreadsheet string, name string, opts ...Option) *MergeRemoteSpreadsheetRequest {
	req := &MergeRemoteSpreadsheetRequest{
		mergedSpreadsheet: mergedSpreadsheet,
		name:              name,
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
	if val, ok := cfg.Params["fontsLocation"].(string); ok {
		req.fontsLocation = val
	}
	if val, ok := cfg.Params["mergeInOneSheet"].(*bool); ok {
		req.mergeInOneSheet = val
	}
	if val, ok := cfg.Params["outFormat"].(string); ok {
		req.outFormat = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
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

func (request *MergeRemoteSpreadsheetRequest) Validate() error {
	if request.mergedSpreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "mergedSpreadsheet")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *MergeRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *MergeRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *MergeRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *MergeRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *MergeRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/merge/spreadsheet"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *MergeRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("mergedSpreadsheet", fmt.Sprintf("%v", request.mergedSpreadsheet))
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.outFormat != "" {
		localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
	}
	if request.mergeInOneSheet != nil {
		localVarQueryParams.Add("mergeInOneSheet", fmt.Sprintf("%v", *request.mergeInOneSheet))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
	if request.fontsLocation != "" {
		localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
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

func (request *MergeRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *MergeRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
