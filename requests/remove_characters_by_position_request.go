package requests

import (
	"fmt"
	"net/url"
	"path/filepath"
)

type RemoveCharactersByPositionRequest struct {
	Spreadsheet     string
	SpreadsheetData []byte
	SpreadsheetName string

	allCharactersAfterText  string
	allCharactersBeforeText string
	caseSensitive           *bool
	outPath                 string
	outStorageName          string
	password                string
	_range                  string
	region                  string
	theFirstNCharacters     *int
	theLastNCharacters      *int
	worksheet               string

	extraQueryParameters map[string]string
}

func NewRemoveCharactersByPositionRequest(Spreadsheet string, opts ...Option) *RemoveCharactersByPositionRequest {
	req := &RemoveCharactersByPositionRequest{
		Spreadsheet: Spreadsheet,
	}
	cfg := &requestConfig{
		Params: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	if val, ok := cfg.Params["allCharactersAfterText"].(string); ok {
		req.allCharactersAfterText = val
	}
	if val, ok := cfg.Params["allCharactersBeforeText"].(string); ok {
		req.allCharactersBeforeText = val
	}
	if val, ok := cfg.Params["caseSensitive"].(*bool); ok {
		req.caseSensitive = val
	}
	if val, ok := cfg.Params["outPath"].(string); ok {
		req.outPath = val
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
	if val, ok := cfg.Params["theFirstNCharacters"].(*int); ok {
		req.theFirstNCharacters = val
	}
	if val, ok := cfg.Params["theLastNCharacters"].(*int); ok {
		req.theLastNCharacters = val
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

func (request *RemoveCharactersByPositionRequest) Validate() error {
	if request.SpreadsheetData == nil && request.Spreadsheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "Spreadsheet")
	}

	return nil
}

func (request *RemoveCharactersByPositionRequest) SetSpreadsheetBytes(data []byte, name string) {
	if name == "" {
		name = "Spreadsheet"
	}
	request.SpreadsheetData = data
	request.SpreadsheetName = name
}

func (request *RemoveCharactersByPositionRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *RemoveCharactersByPositionRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *RemoveCharactersByPositionRequest) GetMethod() string {
	return "PUT"
}

func (request *RemoveCharactersByPositionRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "multipart/form-data"
	return localVarHeaderParams
}

func (request *RemoveCharactersByPositionRequest) GetPath() string {
	localVarPath := "/v4.0/cells/content/remove/characters-by-position"
	return localVarPath
}

func (request *RemoveCharactersByPositionRequest) GetQueryParameters() url.Values {
	localVarQueryParams := url.Values{}
	if request.theFirstNCharacters != nil {
		localVarQueryParams.Add("theFirstNCharacters", fmt.Sprintf("%v", *request.theFirstNCharacters))
	}
	if request.theLastNCharacters != nil {
		localVarQueryParams.Add("theLastNCharacters", fmt.Sprintf("%v", *request.theLastNCharacters))
	}
	if request.allCharactersBeforeText != "" {
		localVarQueryParams.Add("allCharactersBeforeText", fmt.Sprintf("%v", request.allCharactersBeforeText))
	}
	if request.allCharactersAfterText != "" {
		localVarQueryParams.Add("allCharactersAfterText", fmt.Sprintf("%v", request.allCharactersAfterText))
	}
	if request.caseSensitive != nil {
		localVarQueryParams.Add("caseSensitive", fmt.Sprintf("%v", *request.caseSensitive))
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

func (request *RemoveCharactersByPositionRequest) GetJSONBody() interface{} {
	return nil
}

func (request *RemoveCharactersByPositionRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	if request.SpreadsheetData != nil {
		localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
	} else if request.Spreadsheet != "" {
		localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
	}
	return localVarFormParams
}
