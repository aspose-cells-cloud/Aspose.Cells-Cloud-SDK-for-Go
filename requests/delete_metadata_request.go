package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type DeleteMetadataRequest struct {
	File     string
	FileData []byte
	FileName string

	checkExcelRestriction *bool
	outFormat             string
	password              string
	_type                 string

	extraQueryParameters map[string]string
}

func NewDeleteMetadataRequest(File string, opts ...Option) *DeleteMetadataRequest {
	req := &DeleteMetadataRequest{
		File: File,
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
	if val, ok := cfg.Params["outFormat"].(string); ok {
		req.outFormat = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["type"].(string); ok {
		req._type = val
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

func (request *DeleteMetadataRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	return nil
}

func (request *DeleteMetadataRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *DeleteMetadataRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *DeleteMetadataRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *DeleteMetadataRequest) GetMethod() string {
	return "POST"
}

func (request *DeleteMetadataRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *DeleteMetadataRequest) GetPath() string {
	localVarPath := "/v3.0/cells/metadata/delete"
	return localVarPath
}

func (request *DeleteMetadataRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request._type != "" {
		localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
	}
	if request.outFormat != "" {
		localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	if request.checkExcelRestriction != nil {
		localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *DeleteMetadataRequest) GetJSONBody() interface{} {
	return nil
}

func (request *DeleteMetadataRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
