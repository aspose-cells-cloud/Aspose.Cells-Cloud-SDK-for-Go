package requests

import (
	"fmt"
	"net/url"
)

type DecomposeUserTaskRequest struct {
	TaskDescription string

	password string
	region   string

	extraQueryParameters map[string]string
}

func NewDecomposeUserTaskRequest(TaskDescription string, opts ...Option) *DecomposeUserTaskRequest {
	req := &DecomposeUserTaskRequest{
		TaskDescription: TaskDescription,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
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

func (request *DecomposeUserTaskRequest) Validate() error {
	if request.TaskDescription == "" {
		return fmt.Errorf("required request parameter %q is missing", "TaskDescription")
	}

	return nil
}

func (request *DecomposeUserTaskRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *DecomposeUserTaskRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *DecomposeUserTaskRequest) GetMethod() string {
	return "PUT"
}

func (request *DecomposeUserTaskRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *DecomposeUserTaskRequest) GetPath() string {
	localVarPath := "/v4.0/cells/ai/task/decompose"
	return localVarPath
}

func (request *DecomposeUserTaskRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.region != "" {
		localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
	}
	if request.password != "" {
		localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
	}
	for k, v := range request.extraQueryParameters {
		localVarQueryParams.Add(k, v)
	}
	return localVarQueryParams
}

func (request *DecomposeUserTaskRequest) GetJSONBody() interface{} {
	return &request.TaskDescription
}

func (request *DecomposeUserTaskRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
