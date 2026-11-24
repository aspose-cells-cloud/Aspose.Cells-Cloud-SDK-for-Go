package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type PostRepairRequest struct {
	File     string
	FileData []byte
	FileName string

	outFormat string

	extraQueryParameters map[string]string
}

func NewPostRepairRequest(File string, opts ...Option) *PostRepairRequest {
	req := &PostRepairRequest{
		File: File,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["outFormat"].(string); ok {
		req.outFormat = val
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

func (request *PostRepairRequest) Validate() error {
	if request.FileData == nil && request.File == "" {
		return fmt.Errorf("required request parameter %q is missing", "File")
	}

	return nil
}

func (request *PostRepairRequest) SetFileBytes(data []byte, name string) {
	if name == "" {
		name = "File"
	}
	request.FileData = data
	request.FileName = name
}

func (request *PostRepairRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *PostRepairRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *PostRepairRequest) GetMethod() string {
	return "POST"
}

func (request *PostRepairRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *PostRepairRequest) GetPath() string {
	localVarPath := "/v3.0/cells/repair"
	return localVarPath
}

func (request *PostRepairRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.outFormat != "" {
		localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *PostRepairRequest) GetJSONBody() interface{} {
	return nil
}

func (request *PostRepairRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.FileData != nil {
		localVarFormParams[request.FileName] = request.FileData
	} else if request.File != "" {
		localVarFormParams["@"+filepath.Base(request.File)] = request.File
	}
	return localVarFormParams
}
