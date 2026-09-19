package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type StorageExistsRequest struct {
	storageName string

	extraQueryParameters map[string]string
}

func NewStorageExistsRequest(storageName string, opts ...Option) *StorageExistsRequest {
	req := &StorageExistsRequest{
		storageName: storageName,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
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

func (request *StorageExistsRequest) Validate() error {
	if request.storageName == "" {
		return fmt.Errorf("required request parameter %q is missing", "storageName")
	}

	return nil
}

func (request *StorageExistsRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *StorageExistsRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *StorageExistsRequest) GetMethod() string {
	return "GET"
}

func (request *StorageExistsRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *StorageExistsRequest) GetPath() string {
	localVarPath := "/v4.0/cells/storage/{storageName}/exist"
	localVarPath = strings.Replace(localVarPath, "{"+"storageName"+"}", url.PathEscape(fmt.Sprintf("%v", request.storageName)), -1)
	return localVarPath
}

func (request *StorageExistsRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *StorageExistsRequest) GetJSONBody() interface{} {
	return nil
}

func (request *StorageExistsRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
