package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostWorkbookDataCleansingRequest struct {
	dataCleansing *models.DataCleansing
	name          string

	checkExcelRestriction *bool
	folder                string
	password              string
	region                string
	storageName           string

	extraQueryParameters map[string]string
}

func NewPostWorkbookDataCleansingRequest(dataCleansing *models.DataCleansing, name string, opts ...Option) *PostWorkbookDataCleansingRequest {
	req := &PostWorkbookDataCleansingRequest{
		dataCleansing: dataCleansing,
		name:          name,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["checkExcelRestriction"].(*bool); ok {
		req.checkExcelRestriction = val
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

func (request *PostWorkbookDataCleansingRequest) Validate() error {
	if request.dataCleansing == nil {
		return fmt.Errorf("required request parameter %q is missing", "dataCleansing")
	}

	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	return nil
}

func (request *PostWorkbookDataCleansingRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostWorkbookDataCleansingRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostWorkbookDataCleansingRequest) GetMethod() string {
	return "POST"
}

func (request *PostWorkbookDataCleansingRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostWorkbookDataCleansingRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/datacleansing"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PostWorkbookDataCleansingRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.checkExcelRestriction != nil {
		localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostWorkbookDataCleansingRequest) GetJSONBody() interface{} {
	return &request.dataCleansing
}

func (request *PostWorkbookDataCleansingRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
