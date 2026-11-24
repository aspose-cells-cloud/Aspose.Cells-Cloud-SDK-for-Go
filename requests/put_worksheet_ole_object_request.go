package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type PutWorksheetOleObjectRequest struct {
	name      string
	sheetName string

	folder          string
	height          *int
	imageFile       string
	oleFile         string
	storageName     string
	upperLeftColumn *int
	upperLeftRow    *int
	width           *int

	extraQueryParameters map[string]string
}

func NewPutWorksheetOleObjectRequest(name string, sheetName string, opts ...Option) *PutWorksheetOleObjectRequest {
	req := &PutWorksheetOleObjectRequest{
		name:      name,
		sheetName: sheetName,
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
	if val, ok := cfg.Params["height"].(*int); ok {
		req.height = val
	}
	if val, ok := cfg.Params["imageFile"].(string); ok {
		req.imageFile = val
	}
	if val, ok := cfg.Params["oleFile"].(string); ok {
		req.oleFile = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["upperLeftColumn"].(*int); ok {
		req.upperLeftColumn = val
	}
	if val, ok := cfg.Params["upperLeftRow"].(*int); ok {
		req.upperLeftRow = val
	}
	if val, ok := cfg.Params["width"].(*int); ok {
		req.width = val
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

func (request *PutWorksheetOleObjectRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetOleObjectRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetOleObjectRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetOleObjectRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetOleObjectRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetOleObjectRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/oleobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetOleObjectRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.upperLeftRow != nil {
		localVarQueryParams.Add("upperLeftRow", fmt.Sprintf("%v", *request.upperLeftRow))
	}
	if request.upperLeftColumn != nil {
		localVarQueryParams.Add("upperLeftColumn", fmt.Sprintf("%v", *request.upperLeftColumn))
	}
	if request.height != nil {
		localVarQueryParams.Add("height", fmt.Sprintf("%v", *request.height))
	}
	if request.width != nil {
		localVarQueryParams.Add("width", fmt.Sprintf("%v", *request.width))
	}
	if request.oleFile != "" {
		localVarQueryParams.Add("oleFile", fmt.Sprintf("%v", request.oleFile))
	}
	if request.imageFile != "" {
		localVarQueryParams.Add("imageFile", fmt.Sprintf("%v", request.imageFile))
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

func (request *PutWorksheetOleObjectRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PutWorksheetOleObjectRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
