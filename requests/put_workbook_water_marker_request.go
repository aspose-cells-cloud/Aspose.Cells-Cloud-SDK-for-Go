package requests

import (
	"fmt"
	"net/url"
	"strings"

	"asposecellscloud/models"
)

type PutWorkbookWaterMarkerRequest struct {
	name                   string
	textWaterMarkerRequest *models.TextWaterMarkerRequest

	folder      string
	storageName string

	extraQueryParameters map[string]string
}

func NewPutWorkbookWaterMarkerRequest(name string, textWaterMarkerRequest *models.TextWaterMarkerRequest, opts ...Option) *PutWorkbookWaterMarkerRequest {
	req := &PutWorkbookWaterMarkerRequest{
		name:                   name,
		textWaterMarkerRequest: textWaterMarkerRequest,
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

func (request *PutWorkbookWaterMarkerRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request.textWaterMarkerRequest == nil {
		return fmt.Errorf("required request parameter %q is missing", "textWaterMarkerRequest")
	}

	return nil
}

func (request *PutWorkbookWaterMarkerRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PutWorkbookWaterMarkerRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PutWorkbookWaterMarkerRequest) GetMethod() string {
	return "PUT"
}

func (request *PutWorkbookWaterMarkerRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *PutWorkbookWaterMarkerRequest) GetPath() string {
	localVarPath := "/v3.0/cells/{name}/watermarker"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	return localVarPath
}

func (request *PutWorkbookWaterMarkerRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
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

func (request *PutWorkbookWaterMarkerRequest) GetJSONBody() interface{} {
	return &request.textWaterMarkerRequest
}

func (request *PutWorkbookWaterMarkerRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
