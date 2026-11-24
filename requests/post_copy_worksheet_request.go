package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PostCopyWorksheetRequest struct {
	name        string
	options     *models.CopyOptions
	sheetName   string
	sourceSheet string

	folder         string
	sourceFolder   string
	sourceWorkbook string
	storageName    string

	extraQueryParameters map[string]string
}

func NewPostCopyWorksheetRequest(name string, options *models.CopyOptions, sheetName string, sourceSheet string, opts ...Option) *PostCopyWorksheetRequest {
	req := &PostCopyWorksheetRequest{
		name:        name,
		options:     options,
		sheetName:   sheetName,
		sourceSheet: sourceSheet,
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
	if val, ok := cfg.Params["sourceFolder"].(string); ok {
		req.sourceFolder = val
	}
	if val, ok := cfg.Params["sourceWorkbook"].(string); ok {
		req.sourceWorkbook = val
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

func (request *PostCopyWorksheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.options == nil {
		return fmt.Errorf("required request parameter %q is missing", "options")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	if request.sourceSheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "sourceSheet")
	}

	return nil
}

func (request *PostCopyWorksheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostCopyWorksheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostCopyWorksheetRequest) GetMethod() string {
	return "POST"
}

func (request *PostCopyWorksheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PostCopyWorksheetRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PostCopyWorksheetRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("sourceSheet", fmt.Sprintf("%v", request.sourceSheet))
	if request.sourceWorkbook != "" {
		localVarQueryParams.Add("sourceWorkbook", fmt.Sprintf("%v", request.sourceWorkbook))
	}
	if request.sourceFolder != "" {
		localVarQueryParams.Add("sourceFolder", fmt.Sprintf("%v", request.sourceFolder))
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

func (request *PostCopyWorksheetRequest) GetJSONBody() interface{} {
	return &request.options
}

func (request *PostCopyWorksheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
