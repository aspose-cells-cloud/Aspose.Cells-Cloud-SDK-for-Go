package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PutWorksheetAddPictureRequest struct {
	name      string
	sheetName string

	folder           string
	lowerRightColumn *int
	lowerRightRow    *int
	picture          *models.Picture
	picturePath      string
	storageName      string
	upperLeftColumn  *int
	upperLeftRow     *int

	extraQueryParameters map[string]string
}

func NewPutWorksheetAddPictureRequest(name string, sheetName string, opts ...Option) *PutWorksheetAddPictureRequest {
	req := &PutWorksheetAddPictureRequest{
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
	if val, ok := cfg.Params["lowerRightColumn"].(*int); ok {
		req.lowerRightColumn = val
	}
	if val, ok := cfg.Params["lowerRightRow"].(*int); ok {
		req.lowerRightRow = val
	}
	if val, ok := cfg.Params["picture"].(*models.Picture); ok {
		req.picture = val
	}
	if val, ok := cfg.Params["picturePath"].(string); ok {
		req.picturePath = val
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

func (request *PutWorksheetAddPictureRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.sheetName == "" {
		return fmt.Errorf("required request parameter %q is missing", "sheetName")
	}

	return nil
}

func (request *PutWorksheetAddPictureRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorksheetAddPictureRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorksheetAddPictureRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorksheetAddPictureRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorksheetAddPictureRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/worksheets/{sheetName}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", url.PathEscape(fmt.Sprintf("%v", request.sheetName)), -1)
	return localVarPath
}

func (request *PutWorksheetAddPictureRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.upperLeftRow != nil {
		localVarQueryParams.Add("upperLeftRow", fmt.Sprintf("%v", *request.upperLeftRow))
	}
	if request.upperLeftColumn != nil {
		localVarQueryParams.Add("upperLeftColumn", fmt.Sprintf("%v", *request.upperLeftColumn))
	}
	if request.lowerRightRow != nil {
		localVarQueryParams.Add("lowerRightRow", fmt.Sprintf("%v", *request.lowerRightRow))
	}
	if request.lowerRightColumn != nil {
		localVarQueryParams.Add("lowerRightColumn", fmt.Sprintf("%v", *request.lowerRightColumn))
	}
	if request.picturePath != "" {
		localVarQueryParams.Add("picturePath", fmt.Sprintf("%v", request.picturePath))
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

func (request *PutWorksheetAddPictureRequest) GetJSONBody() interface{} {
	return &request.picture
}

func (request *PutWorksheetAddPictureRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
