package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type CopyFolderRequest struct {
	destPath string
	srcPath  string

	destStorageName string
	srcStorageName  string

	extraQueryParameters map[string]string
}

func NewCopyFolderRequest(destPath string, srcPath string, opts ...Option) *CopyFolderRequest {
	req := &CopyFolderRequest{
		destPath: destPath,
		srcPath:  srcPath,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["destStorageName"].(string); ok {
		req.destStorageName = val
	}
	if val, ok := cfg.Params["srcStorageName"].(string); ok {
		req.srcStorageName = val
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

func (request *CopyFolderRequest) Validate() error {
	if request.destPath == "" {
		return fmt.Errorf("required request parameter %q is missing", "destPath")
	}

	if request.srcPath == "" {
		return fmt.Errorf("required request parameter %q is missing", "srcPath")
	}

	return nil
}

func (request *CopyFolderRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *CopyFolderRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *CopyFolderRequest) GetMethod() string {
	return "PUT"
}

func (request *CopyFolderRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *CopyFolderRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/folder/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", url.PathEscape(fmt.Sprintf("%v", request.srcPath)), -1)
	return localVarPath
}

func (request *CopyFolderRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("destPath", fmt.Sprintf("%v", request.destPath))
	if request.srcStorageName != "" {
		localVarQueryParams.Add("srcStorageName", fmt.Sprintf("%v", request.srcStorageName))
	}
	if request.destStorageName != "" {
		localVarQueryParams.Add("destStorageName", fmt.Sprintf("%v", request.destStorageName))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *CopyFolderRequest) GetJSONBody() interface{} {
	return nil
}

func (request *CopyFolderRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
