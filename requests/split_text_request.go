package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type SplitTextRequest struct {
	delimiters      string
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	HowToSplit                     string
	keepDelimitersInResultingCells *bool
	keepDelimitersPosition         string
	outPath                        string
	outPositionRange               string
	outStorageName                 string
	password                       string
	_range                         string
	region                         string
	worksheet                      string

	extraQueryParameters map[string]string
}

func NewSplitTextRequest(delimiters string, Spreadsheet string, opts ...Option) *SplitTextRequest {
	req := &SplitTextRequest{
		delimiters:  delimiters,
		Spreadsheet: Spreadsheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["HowToSplit"].(string); ok {
		req.HowToSplit = val
	}
	if val, ok := cfg.Params["keepDelimitersInResultingCells"].(*bool); ok {
		req.keepDelimitersInResultingCells = val
	}
	if val, ok := cfg.Params["keepDelimitersPosition"].(string); ok {
		req.keepDelimitersPosition = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
	}
	if val, ok := cfg.Params["outPositionRange"].(string); ok {
		req.outPositionRange = val
	}
	if val, ok := cfg.Params["outStorageName"].(string); ok {
		req.outStorageName = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["range"].(string); ok {
		req._range = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["worksheet"].(string); ok {
		req.worksheet = val
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

func (request *SplitTextRequest) Validate() error {
	if request.delimiters == "" {
		return fmt.Errorf("required request parameter %q is missing", "delimiters")
	}

	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *SplitTextRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *SplitTextRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *SplitTextRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *SplitTextRequest) GetMethod() string {
	return "PUT"
}

func (request *SplitTextRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *SplitTextRequest) GetPath() string {
	localVarPath := "/v4.0/cells/content/split/text"
	return localVarPath
}

func (request *SplitTextRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("delimiters", fmt.Sprintf("%v", request.delimiters))
	if request.keepDelimitersInResultingCells != nil {
		localVarQueryParams.Add("keepDelimitersInResultingCells", fmt.Sprintf("%v", *request.keepDelimitersInResultingCells))
	}
	if request.keepDelimitersPosition != "" {
		localVarQueryParams.Add("keepDelimitersPosition", fmt.Sprintf("%v", request.keepDelimitersPosition))
	}
	if request.HowToSplit != "" {
		localVarQueryParams.Add("HowToSplit", fmt.Sprintf("%v", request.HowToSplit))
	}
	if request.outPositionRange != "" {
		localVarQueryParams.Add("outPositionRange", fmt.Sprintf("%v", request.outPositionRange))
	}
	if request.worksheet != "" {
		localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
	}
	if request._range != "" {
		localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
	}
	if request.outPath != "" {
		localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
	}
	if request.outStorageName != "" {
		localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
	}
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

func (request *SplitTextRequest) GetJSONBody() interface{} {
	return nil
}

func (request *SplitTextRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
