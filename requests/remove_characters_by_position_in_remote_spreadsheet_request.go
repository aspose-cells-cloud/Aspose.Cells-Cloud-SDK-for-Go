package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type RemoveCharactersByPositionInRemoteSpreadsheetRequest struct {
	name      string
	_range    string
	worksheet string

	allCharactersAfterText  string
	allCharactersBeforeText string
	caseSensitive           *bool
	folder                  string
	password                string
	region                  string
	storageName             string
	theFirstNCharacters     *int
	theLastNCharacters      *int

	extraQueryParameters map[string]string
}

func NewRemoveCharactersByPositionInRemoteSpreadsheetRequest(name string, _range string, worksheet string, opts ...Option) *RemoveCharactersByPositionInRemoteSpreadsheetRequest {
	req := &RemoveCharactersByPositionInRemoteSpreadsheetRequest{
		name:      name,
		_range:    _range,
		worksheet: worksheet,
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
	if val, ok := cfg.Params["folder"].(string); ok {
		req.folder = val
	}
	if val, ok := cfg.Params["password"].(string); ok {
		req.password = val
	}
	if val, ok := cfg.Params["region"].(string); ok {
		req.region = val
	}
	if val, ok := cfg.Params["storageName"].(string); ok {
		req.storageName = val
	}
	if val, ok := cfg.Params["theFirstNCharacters"].(*int); ok {
		req.theFirstNCharacters = val
	}
	if val, ok := cfg.Params["theLastNCharacters"].(*int); ok {
		req.theLastNCharacters = val
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

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) Validate() error {
	if request.name == "" {
		return fmt.Errorf("required request parameter %q is missing", "name")
	}

	if request._range == "" {
		return fmt.Errorf("required request parameter %q is missing", "range")
	}

	if request.worksheet == "" {
		return fmt.Errorf("required request parameter %q is missing", "worksheet")
	}

	return nil
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	request.extraQueryParameters[key] = value
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
	if request.extraQueryParameters == nil {
		request.extraQueryParameters = make(map[string]string)
	}
	for k, v := range params {
		request.extraQueryParameters[k] = v
	}
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) GetMethod() string {
	return "PUT"
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"
	return localVarHeaderParams
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) GetPath() string {
	localVarPath := "/v4.0/cells/{name}/worksheets/{worksheet}/range/{range}/content/remove/characters-by-position"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(fmt.Sprintf("%v", request.name)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", url.PathEscape(fmt.Sprintf("%v", request.worksheet)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", url.PathEscape(fmt.Sprintf("%v", request._range)), -1)
	return localVarPath
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
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
	if request.folder != "" {
		localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
	}
	if request.storageName != "" {
		localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
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

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
	return nil
}

func (request *RemoveCharactersByPositionInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
	localVarFormParams := make(map[string]interface{})
	return localVarFormParams
}
