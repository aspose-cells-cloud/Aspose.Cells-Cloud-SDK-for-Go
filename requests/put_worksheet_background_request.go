package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

type PutWorksheetBackgroundRequest struct {
	name      string
	sheetName string

	File             string
	FileData         []byte
	FileName         string
	folder           string
	imageAdaptOption string
	picPath          string
	storageName      string

	extraQueryParameters map[string]string
}

func NewPutWorksheetBackgroundRequest(name string, sheetName string, opts ...Option) *PutWorksheetBackgroundRequest {
	req := &PutWorksheetBackgroundRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["File"].(string); ok {
		req.File = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["imageAdaptOption"].(string); ok {
		req.imageAdaptOption = val
	}
	if val, ok := cfg.Params["picPath"].(string); ok {
		req.picPath = val
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

func (request *PutWorksheetBackgroundRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetBackgroundRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PutWorksheetBackgroundRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetBackgroundRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetBackgroundRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetBackgroundRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PutWorksheetBackgroundRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/background"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetBackgroundRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.picPath != "" {
		localVarQueryParams.Add("picPath", fmt.Sprintf("%v", request.picPath))
	}
	if request.imageAdaptOption != "" {
		localVarQueryParams.Add("imageAdaptOption", fmt.Sprintf("%v", request.imageAdaptOption))
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

func (request *PutWorksheetBackgroundRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetBackgroundRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
