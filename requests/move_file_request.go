package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type MoveFileRequest struct {
	destPath string
	srcPath  string

	destStorageName string
	srcStorageName  string
	versionId       string

	extraQueryParameters map[string]string
}

func NewMoveFileRequest(destPath string, srcPath string, opts ...Option) *MoveFileRequest {
	req := &MoveFileRequest{
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
	if val, ok := cfg.Params["versionId"].(string); ok {
		req.versionId = val
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

func (request *MoveFileRequest) Validate() error {
	if request.destPath == "" {
		return fmt.Errorf("required request parameter %q is missing", "destPath")
	}

	if request.srcPath == "" {
		return fmt.Errorf("required request parameter %q is missing", "srcPath")
	}

	return nil
}

func (request *MoveFileRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *MoveFileRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *MoveFileRequest) GetMethod() string {
	return "PUT"
}

func (request *MoveFileRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *MoveFileRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/file/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", url.PathEscape(fmt.Sprintf("%v", request.srcPath)), -1)
	return localVarPath
}

func (request *MoveFileRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("destPath", fmt.Sprintf("%v", request.destPath))
	if request.srcStorageName != "" {
		localVarQueryParams.Add("srcStorageName", fmt.Sprintf("%v", request.srcStorageName))
	}
	if request.destStorageName != "" {
		localVarQueryParams.Add("destStorageName", fmt.Sprintf("%v", request.destStorageName))
	}
	if request.versionId != "" {
		localVarQueryParams.Add("versionId", fmt.Sprintf("%v", request.versionId))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *MoveFileRequest) GetJSONBody() interface{} {
	return nil
}

func (request *MoveFileRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
