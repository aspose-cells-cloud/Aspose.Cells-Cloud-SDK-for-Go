package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PutWorksheetShapeRequest struct {
	name      string
	sheetName string

	DrawingType     string
	folder          string
	height          *int
	left            *int
	shapeDTO        models.ShapeLike
	storageName     string
	top             *int
	upperLeftColumn *int
	upperLeftRow    *int
	width           *int

	extraQueryParameters map[string]string
}

func NewPutWorksheetShapeRequest(name string, sheetName string, opts ...Option) *PutWorksheetShapeRequest {
	req := &PutWorksheetShapeRequest{
		name:      name,
		sheetName: sheetName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["DrawingType"].(string); ok {
		req.DrawingType = val
	}
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["height"].(*int); ok {
		req.height = val
	}
	if val, ok := cfg.Params["left"].(*int); ok {
		req.left = val
	}
	if val, ok := cfg.Params["shapeDTO"].(models.ShapeLike); ok {
		req.shapeDTO = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["top"].(*int); ok {
		req.top = val
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

func (request *PutWorksheetShapeRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetShapeRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetShapeRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetShapeRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetShapeRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetShapeRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/shapes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetShapeRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.DrawingType != "" {
		localVarQueryParams.Add("DrawingType", fmt.Sprintf("%v", request.DrawingType))
	}
	if request.upperLeftRow != nil {
		localVarQueryParams.Add("upperLeftRow", fmt.Sprintf("%v", *request.upperLeftRow))
	}
	if request.upperLeftColumn != nil {
		localVarQueryParams.Add("upperLeftColumn", fmt.Sprintf("%v", *request.upperLeftColumn))
	}
	if request.top != nil {
		localVarQueryParams.Add("top", fmt.Sprintf("%v", *request.top))
	}
	if request.left != nil {
		localVarQueryParams.Add("left", fmt.Sprintf("%v", *request.left))
	}
	if request.width != nil {
		localVarQueryParams.Add("width", fmt.Sprintf("%v", *request.width))
	}
	if request.height != nil {
		localVarQueryParams.Add("height", fmt.Sprintf("%v", *request.height))
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

func (request *PutWorksheetShapeRequest) GetJSONBody() interface{} {
	return &request.shapeDTO
}

func (request *PutWorksheetShapeRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
