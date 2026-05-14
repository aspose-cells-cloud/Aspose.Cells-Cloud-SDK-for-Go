/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="api_cells">
*   Copyright (c) 2026 Aspose.Cells Cloud
* </copyright>
* <summary>
*   Permission is hereby granted, free of charge, to any person obtaining a copy
*  of this software and associated documentation files (the "Software"), to deal
*  in the Software without restriction, including without limitation the rights
*  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  copies of the Software, and to permit persons to whom the Software is
*  furnished to do so, subject to the following conditions:
* 
*  The above copyright notice and this permission notice shall be included in all
*  copies or substantial portions of the Software.
* 
*  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*  SOFTWARE.
* </summary> 
-------------------------------------------------------------------------------------------------------------------- **/

package asposecellscloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
type CellsCloudOption struct {
	OptionName  string
	OptionValue string
}

func Version() {
	fmt.Println("---Version: 26.5---")
}

func NewCellsApiService(appSid string, appKey string, opts ...string) *CellsApiService {
	var basePath = ""
	var version = ""
	for i, v := range opts {
		switch i {
		case 0:
			basePath = v
		case 1:
			version = v
		}
	}
	config := NewConfiguration(appSid, appKey, basePath, version)
	client := NewAPIClient(config)
	return client.CellsApi
}

type CellsApiService service

func (a *CellsApiService) DecomposeUserTask(data *DecomposeUserTaskRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) TranslateSpreadsheet(data *TranslateSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) TranslateTextFile(data *TranslateTextFileRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) AggregateCellsByColor(data *AggregateCellsByColorRequest  ) (  AggregateResultByColorResponse,  *http.Response, error) {
	var (
	localVarReturnValue AggregateResultByColorResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) MathCalculate(data *MathCalculateRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAccessToken(data *PostAccessTokenRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetAsposeCellsCloudStatus(data *GetAsposeCellsCloudStatusRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CheckCloudServiceHealth(data *CheckCloudServiceHealthRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ExportSpreadsheetAsFormat(data *ExportSpreadsheetAsFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ExportWorksheetAsFormat(data *ExportWorksheetAsFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ExportChartAsFormat(data *ExportChartAsFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ExportTableAsFormat(data *ExportTableAsFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ExportRangeAsFormat(data *ExportRangeAsFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertSpreadsheet(data *ConvertSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertSpreadsheetToPdf(data *ConvertSpreadsheetToPdfRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertSpreadsheetToJson(data *ConvertSpreadsheetToJsonRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertSpreadsheetToCsv(data *ConvertSpreadsheetToCsvRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertWorksheetToImage(data *ConvertWorksheetToImageRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertWorksheetToPdf(data *ConvertWorksheetToPdfRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertWorksheetToJson(data *ConvertWorksheetToJsonRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertWorksheetToCsv(data *ConvertWorksheetToCsvRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertWorksheetToHtml(data *ConvertWorksheetToHtmlRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertWorksheetToHtmlTable(data *ConvertWorksheetToHtmlTableRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertTableToImage(data *ConvertTableToImageRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertTableToPdf(data *ConvertTableToPdfRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertTableToCsv(data *ConvertTableToCsvRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertTableToHtml(data *ConvertTableToHtmlRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertTableToJson(data *ConvertTableToJsonRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertRangeToImage(data *ConvertRangeToImageRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertRangeToPdf(data *ConvertRangeToPdfRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertRangeToCsv(data *ConvertRangeToCsvRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertRangeToHtml(data *ConvertRangeToHtmlRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertRangeToJson(data *ConvertRangeToJsonRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertChartToImage(data *ConvertChartToImageRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertChartToPdf(data *ConvertChartToPdfRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SaveSpreadsheetAs(data *SaveSpreadsheetAsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) MergeSpreadsheets(data *MergeSpreadsheetsRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) MergeSpreadsheetsInRemoteFolder(data *MergeSpreadsheetsInRemoteFolderRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) MergeRemoteSpreadsheet(data *MergeRemoteSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SplitSpreadsheet(data *SplitSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SplitTable(data *SplitTableRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SplitRemoteSpreadsheet(data *SplitRemoteSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ImportDataIntoSpreadsheet(data *ImportDataIntoSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ImportJSONDataIntoSpreadsheet(data *ImportJSONDataIntoSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ImportXMLDataIntoSpreadsheet(data *ImportXMLDataIntoSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ImportCSVDataIntoSpreadsheet(data *ImportCSVDataIntoSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetPublicKey(data *GetPublicKeyRequest  ) (  CellsCloudPublicKeyResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudPublicKeyResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CreateSpreadsheet(data *CreateSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetsWithLocalSpreadsheet(data *GetWorksheetsWithLocalSpreadsheetRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) AddWorksheetToSpreadsheet(data *AddWorksheetToSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetFromSpreadsheet(data *DeleteWorksheetFromSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RenameWorksheetInSpreadsheet(data *RenameWorksheetInSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) MoveWorksheetInSpreadsheet(data *MoveWorksheetInSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CompressSpreadsheet(data *CompressSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RepairSpreadsheet(data *RepairSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetMergedCellsInRemotedWorksheet(data *GetMergedCellsInRemotedWorksheetRequest  ) (  []CellArea,  *http.Response, error) {
	var (
	localVarReturnValue []CellArea 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetMergedCellsInWorksheet(data *GetMergedCellsInWorksheetRequest  ) (  []CellArea,  *http.Response, error) {
	var (
	localVarReturnValue []CellArea 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) AcceptAllRevisions(data *AcceptAllRevisionsRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) AcceptAllRevisionsInRemoteSpreadsheet(data *AcceptAllRevisionsInRemoteSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetSpreadsheetStructure(data *GetSpreadsheetStructureRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetStructureInRemoteSpreadsheet(data *GetStructureInRemoteSpreadsheetRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ProtectSpreadsheet(data *ProtectSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) UnprotectSpreadsheet(data *UnprotectSpreadsheetRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SpreadsheetDigitalsignature(data *SpreadsheetDigitalsignatureRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchAllTextItemsInRemoteSpreadsheet(data *SearchAllTextItemsInRemoteSpreadsheetRequest  ) (  SearchResponse,  *http.Response, error) {
	var (
	localVarReturnValue SearchResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchSpreadsheetAllTextItems(data *SearchSpreadsheetAllTextItemsRequest  ) (  SearchResponse,  *http.Response, error) {
	var (
	localVarReturnValue SearchResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchSpreadsheetContent(data *SearchSpreadsheetContentRequest  ) (  SearchResponse,  *http.Response, error) {
	var (
	localVarReturnValue SearchResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchContentInRemoteSpreadsheet(data *SearchContentInRemoteSpreadsheetRequest  ) (  SearchResponse,  *http.Response, error) {
	var (
	localVarReturnValue SearchResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchContentInRemoteWorksheet(data *SearchContentInRemoteWorksheetRequest  ) (  SearchResponse,  *http.Response, error) {
	var (
	localVarReturnValue SearchResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchContentInRemoteRange(data *SearchContentInRemoteRangeRequest  ) (  SearchResponse,  *http.Response, error) {
	var (
	localVarReturnValue SearchResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ReplaceSpreadsheetContent(data *ReplaceSpreadsheetContentRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ReplaceContentInRemoteSpreadsheet(data *ReplaceContentInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ReplaceContentInRemoteWorksheet(data *ReplaceContentInRemoteWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ReplaceContentInRemoteRange(data *ReplaceContentInRemoteRangeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchSpreadsheetBrokenLinks(data *SearchSpreadsheetBrokenLinksRequest  ) (  BrokenLinksResponse,  *http.Response, error) {
	var (
	localVarReturnValue BrokenLinksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchBrokenLinksInRemoteSpreadsheet(data *SearchBrokenLinksInRemoteSpreadsheetRequest  ) (  BrokenLinksResponse,  *http.Response, error) {
	var (
	localVarReturnValue BrokenLinksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchBrokenLinksInRemoteWorksheet(data *SearchBrokenLinksInRemoteWorksheetRequest  ) (  BrokenLinksResponse,  *http.Response, error) {
	var (
	localVarReturnValue BrokenLinksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SearchBrokenLinksInRemoteRange(data *SearchBrokenLinksInRemoteRangeRequest  ) (  BrokenLinksResponse,  *http.Response, error) {
	var (
	localVarReturnValue BrokenLinksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) Spec(data *SpecRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) CodegenSpec(data *CodegenSpecRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) TrimCharacterInRemoteSpreadsheet(data *TrimCharacterInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) TrimCharacter(data *TrimCharacterRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) UpdateWordCase(data *UpdateWordCaseRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) UpdateWordCaseInRemoteSpreadsheet(data *UpdateWordCaseInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveCharacters(data *RemoveCharactersRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveCharactersInRemoteSpreadsheet(data *RemoveCharactersInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveCharactersByPosition(data *RemoveCharactersByPositionRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveCharactersByPositionInRemoteSpreadsheet(data *RemoveCharactersByPositionInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveDuplicateSubstrings(data *RemoveDuplicateSubstringsRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveDuplicateSubstringsInRemoteSpreadsheet(data *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) AddText(data *AddTextRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) AddTextInRemoteSpreadsheet(data *AddTextInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertText(data *ConvertTextRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ConvertTextInRemoteSpreadsheet(data *ConvertTextInRemoteSpreadsheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ExtractText(data *ExtractTextRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SplitText(data *SplitTextRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveSpreadsheetBlankRows(data *RemoveSpreadsheetBlankRowsRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveSpreadsheetBlankColumns(data *RemoveSpreadsheetBlankColumnsRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveSpreadsheetBlankWorksheets(data *RemoveSpreadsheetBlankWorksheetsRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) RemoveDuplicates(data *RemoveDuplicatesRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) SwapRange(data *SwapRangeRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) FlipData(data *FlipDataRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) TransposeData(data *TransposeDataRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) UnpivotRange(data *UnpivotRangeRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) UnpivotTable(data *UnpivotTableRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DownloadFile(data *DownloadFileRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) UploadFile(data *UploadFileRequest  ) (  FilesUploadResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesUploadResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CopyFile(data *CopyFileRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) MoveFile(data *MoveFileRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) DeleteFile(data *DeleteFileRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) GetFilesList(data *GetFilesListRequest  ) (  FilesList,  *http.Response, error) {
	var (
	localVarReturnValue FilesList 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CreateFolder(data *CreateFolderRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) CopyFolder(data *CopyFolderRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) MoveFolder(data *MoveFolderRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) DeleteFolder(data *DeleteFolderRequest  ) (   *http.Response, error) {
	var (
	 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return    nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return   localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return    localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
 
	return   localVarHttpResponse, err
}


func (a *CellsApiService) StorageExists(data *StorageExistsRequest  ) (  StorageExist,  *http.Response, error) {
	var (
	localVarReturnValue StorageExist 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) ObjectExists(data *ObjectExistsRequest  ) (  ObjectExist,  *http.Response, error) {
	var (
	localVarReturnValue ObjectExist 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetDiscUsage(data *GetDiscUsageRequest  ) (  DiscUsage,  *http.Response, error) {
	var (
	localVarReturnValue DiscUsage 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetFileVersions(data *GetFileVersionsRequest  ) (  FileVersions,  *http.Response, error) {
	var (
	localVarReturnValue FileVersions 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAnalyzeExcel(data *PostAnalyzeExcelRequest  ) (  []AnalyzedResult,  *http.Response, error) {
	var (
	localVarReturnValue []AnalyzedResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetAutoFilter(data *GetWorksheetAutoFilterRequest  ) (  AutoFilterResponse,  *http.Response, error) {
	var (
	localVarReturnValue AutoFilterResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetDateFilter(data *PutWorksheetDateFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetFilter(data *PutWorksheetFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetIconFilter(data *PutWorksheetIconFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetCustomFilter(data *PutWorksheetCustomFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetDynamicFilter(data *PutWorksheetDynamicFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetFilterTop10(data *PutWorksheetFilterTop10Request  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetColorFilter(data *PutWorksheetColorFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetMatchBlanks(data *PostWorksheetMatchBlanksRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetMatchNonBlanks(data *PostWorksheetMatchNonBlanksRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetAutoFilterRefresh(data *PostWorksheetAutoFilterRefreshRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetDateFilter(data *DeleteWorksheetDateFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetFilter(data *DeleteWorksheetFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetAutoshapes(data *GetWorksheetAutoshapesRequest  ) (  AutoShapesResponse,  *http.Response, error) {
	var (
	localVarReturnValue AutoShapesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetAutoshapeWithFormat(data *GetWorksheetAutoshapeWithFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostBatchConvert(data *PostBatchConvertRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostBatchProtect(data *PostBatchProtectRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostBatchLock(data *PostBatchLockRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostBatchUnlock(data *PostBatchUnlockRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostBatchSplit(data *PostBatchSplitRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostClearContents(data *PostClearContentsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostClearFormats(data *PostClearFormatsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWorksheetRangeStyle(data *PostUpdateWorksheetRangeStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetMerge(data *PostWorksheetMergeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetUnmerge(data *PostWorksheetUnmergeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetCells(data *GetWorksheetCellsRequest  ) (  CellsResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetCell(data *GetWorksheetCellRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetCellStyle(data *GetWorksheetCellStyleRequest  ) (  StyleResponse,  *http.Response, error) {
	var (
	localVarReturnValue StyleResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellSetValue(data *PostWorksheetCellSetValueRequest  ) (  CellResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWorksheetCellStyle(data *PostUpdateWorksheetCellStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSetCellRangeValue(data *PostSetCellRangeValueRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCopyCellIntoCell(data *PostCopyCellIntoCellRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetCellHtmlString(data *GetCellHtmlStringRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSetCellHtmlString(data *PostSetCellHtmlStringRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCellCalculate(data *PostCellCalculateRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCellCharacters(data *PostCellCharactersRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetColumns(data *GetWorksheetColumnsRequest  ) (  ColumnsResponse,  *http.Response, error) {
	var (
	localVarReturnValue ColumnsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSetWorksheetColumnWidth(data *PostSetWorksheetColumnWidthRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetColumn(data *GetWorksheetColumnRequest  ) (  ColumnResponse,  *http.Response, error) {
	var (
	localVarReturnValue ColumnResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutInsertWorksheetColumns(data *PutInsertWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetColumns(data *DeleteWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostHideWorksheetColumns(data *PostHideWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUnhideWorksheetColumns(data *PostUnhideWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostGroupWorksheetColumns(data *PostGroupWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUngroupWorksheetColumns(data *PostUngroupWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCopyWorksheetColumns(data *PostCopyWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostColumnStyle(data *PostColumnStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetRows(data *GetWorksheetRowsRequest  ) (  RowsResponse,  *http.Response, error) {
	var (
	localVarReturnValue RowsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetRow(data *GetWorksheetRowRequest  ) (  RowResponse,  *http.Response, error) {
	var (
	localVarReturnValue RowResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetRow(data *DeleteWorksheetRowRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetRows(data *DeleteWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutInsertWorksheetRows(data *PutInsertWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutInsertWorksheetRow(data *PutInsertWorksheetRowRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWorksheetRow(data *PostUpdateWorksheetRowRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostHideWorksheetRows(data *PostHideWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUnhideWorksheetRows(data *PostUnhideWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostGroupWorksheetRows(data *PostGroupWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUngroupWorksheetRows(data *PostUngroupWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCopyWorksheetRows(data *PostCopyWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRowStyle(data *PostRowStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetCellsCloudServicesHealthCheck(data *GetCellsCloudServicesHealthCheckRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetCellsCloudServiceStatus(data *GetCellsCloudServiceStatusRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartArea(data *GetChartAreaRequest  ) (  ChartAreaResponse,  *http.Response, error) {
	var (
	localVarReturnValue ChartAreaResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartAreaFillFormat(data *GetChartAreaFillFormatRequest  ) (  FillFormatResponse,  *http.Response, error) {
	var (
	localVarReturnValue FillFormatResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartAreaBorder(data *GetChartAreaBorderRequest  ) (  LineResponse,  *http.Response, error) {
	var (
	localVarReturnValue LineResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetCharts(data *GetWorksheetChartsRequest  ) (  ChartsResponse,  *http.Response, error) {
	var (
	localVarReturnValue ChartsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetChart(data *GetWorksheetChartRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetChart(data *PutWorksheetChartRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetChart(data *DeleteWorksheetChartRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetChart(data *PostWorksheetChartRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetChartLegend(data *GetWorksheetChartLegendRequest  ) (  LegendResponse,  *http.Response, error) {
	var (
	localVarReturnValue LegendResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetChartLegend(data *PostWorksheetChartLegendRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetChartLegend(data *PutWorksheetChartLegendRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetChartLegend(data *DeleteWorksheetChartLegendRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetCharts(data *DeleteWorksheetChartsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetChartTitle(data *GetWorksheetChartTitleRequest  ) (  TitleResponse,  *http.Response, error) {
	var (
	localVarReturnValue TitleResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetChartTitle(data *PostWorksheetChartTitleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetChartTitle(data *PutWorksheetChartTitleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetChartTitle(data *DeleteWorksheetChartTitleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartSeriesAxis(data *GetChartSeriesAxisRequest  ) (  AxisResponse,  *http.Response, error) {
	var (
	localVarReturnValue AxisResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartCategoryAxis(data *GetChartCategoryAxisRequest  ) (  AxisResponse,  *http.Response, error) {
	var (
	localVarReturnValue AxisResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartValueAxis(data *GetChartValueAxisRequest  ) (  AxisResponse,  *http.Response, error) {
	var (
	localVarReturnValue AxisResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartSecondCategoryAxis(data *GetChartSecondCategoryAxisRequest  ) (  AxisResponse,  *http.Response, error) {
	var (
	localVarReturnValue AxisResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetChartSecondValueAxis(data *GetChartSecondValueAxisRequest  ) (  AxisResponse,  *http.Response, error) {
	var (
	localVarReturnValue AxisResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostChartSeriesAxis(data *PostChartSeriesAxisRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostChartCategoryAxis(data *PostChartCategoryAxisRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostChartValueAxis(data *PostChartValueAxisRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostChartSecondCategoryAxis(data *PostChartSecondCategoryAxisRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostChartSecondValueAxis(data *PostChartSecondValueAxisRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetConditionalFormattings(data *GetWorksheetConditionalFormattingsRequest  ) (  ConditionalFormattingsResponse,  *http.Response, error) {
	var (
	localVarReturnValue ConditionalFormattingsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetConditionalFormatting(data *GetWorksheetConditionalFormattingRequest  ) (  ConditionalFormattingResponse,  *http.Response, error) {
	var (
	localVarReturnValue ConditionalFormattingResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetConditionalFormatting(data *PutWorksheetConditionalFormattingRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetFormatCondition(data *PutWorksheetFormatConditionRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetFormatConditionArea(data *PutWorksheetFormatConditionAreaRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetFormatConditionCondition(data *PutWorksheetFormatConditionConditionRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetConditionalFormattings(data *DeleteWorksheetConditionalFormattingsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetConditionalFormatting(data *DeleteWorksheetConditionalFormattingRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetConditionalFormattingArea(data *DeleteWorksheetConditionalFormattingAreaRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbook(data *GetWorkbookRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutConvertWorkbook(data *PutConvertWorkbookRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookSaveAs(data *PostWorkbookSaveAsRequest  ) (  SaveResponse,  *http.Response, error) {
	var (
	localVarReturnValue SaveResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToPDF(data *PostConvertWorkbookToPDFRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToPNG(data *PostConvertWorkbookToPNGRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToDocx(data *PostConvertWorkbookToDocxRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToPptx(data *PostConvertWorkbookToPptxRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToHtml(data *PostConvertWorkbookToHtmlRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToMarkdown(data *PostConvertWorkbookToMarkdownRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToJson(data *PostConvertWorkbookToJsonRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToSQL(data *PostConvertWorkbookToSQLRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbookToCSV(data *PostConvertWorkbookToCSVRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorksheetToImage(data *PostConvertWorksheetToImageRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertWorkbook(data *PostConvertWorkbookRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CheckWrokbookExternalReference(data *CheckWrokbookExternalReferenceRequest  ) (  CheckedExternalReferenceResponse,  *http.Response, error) {
	var (
	localVarReturnValue CheckedExternalReferenceResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) CheckWorkbookFormulaErrors(data *CheckWorkbookFormulaErrorsRequest  ) (  CheckedFormulaErrorsResponse,  *http.Response, error) {
	var (
	localVarReturnValue CheckedFormulaErrorsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostExport(data *PostExportRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookExportXML(data *PostWorkbookExportXMLRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookImportJson(data *PostWorkbookImportJsonRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookImportXML(data *PostWorkbookImportXMLRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostImportData(data *PostImportDataRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookDataCleansing(data *PostWorkbookDataCleansingRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostDataCleansing(data *PostDataCleansingRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookDataDeduplication(data *PostWorkbookDataDeduplicationRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostDataDeduplication(data *PostDataDeduplicationRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookDataFill(data *PostWorkbookDataFillRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostDataFill(data *PostDataFillRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostDeleteIncompleteRows(data *PostDeleteIncompleteRowsRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostDataTransformation(data *PostDataTransformationRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetHyperlinks(data *GetWorksheetHyperlinksRequest  ) (  HyperlinksResponse,  *http.Response, error) {
	var (
	localVarReturnValue HyperlinksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetHyperlink(data *GetWorksheetHyperlinkRequest  ) (  HyperlinkResponse,  *http.Response, error) {
	var (
	localVarReturnValue HyperlinkResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetHyperlink(data *DeleteWorksheetHyperlinkRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetHyperlink(data *PostWorksheetHyperlinkRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetHyperlink(data *PutWorksheetHyperlinkRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetHyperlinks(data *DeleteWorksheetHyperlinksRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAssemble(data *PostAssembleRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCompress(data *PostCompressRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostMerge(data *PostMergeRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSplit(data *PostSplitRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSearch(data *PostSearchRequest  ) (  []TextItem,  *http.Response, error) {
	var (
	localVarReturnValue []TextItem 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostReplace(data *PostReplaceRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostImport(data *PostImportRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWatermark(data *PostWatermarkRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostClearObjects(data *PostClearObjectsRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostReverse(data *PostReverseRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRepair(data *PostRepairRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRotate(data *PostRotateRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostMetadata(data *PostMetadataRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetMetadata(data *GetMetadataRequest  ) (  []CellsDocumentProperty,  *http.Response, error) {
	var (
	localVarReturnValue []CellsDocumentProperty 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteMetadata(data *DeleteMetadataRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetListObjects(data *GetWorksheetListObjectsRequest  ) (  ListObjectsResponse,  *http.Response, error) {
	var (
	localVarReturnValue ListObjectsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetListObject(data *GetWorksheetListObjectRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetListObject(data *PutWorksheetListObjectRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetListObjects(data *DeleteWorksheetListObjectsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetListObject(data *DeleteWorksheetListObjectRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListObject(data *PostWorksheetListObjectRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListObjectConvertToRange(data *PostWorksheetListObjectConvertToRangeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListObjectSummarizeWithPivotTable(data *PostWorksheetListObjectSummarizeWithPivotTableRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListObjectSortTable(data *PostWorksheetListObjectSortTableRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListObjectRemoveDuplicates(data *PostWorksheetListObjectRemoveDuplicatesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListObjectInsertSlicer(data *PostWorksheetListObjectInsertSlicerRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListColumn(data *PostWorksheetListColumnRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetListColumnsTotal(data *PostWorksheetListColumnsTotalRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetOleObjects(data *GetWorksheetOleObjectsRequest  ) (  OleObjectsResponse,  *http.Response, error) {
	var (
	localVarReturnValue OleObjectsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetOleObject(data *GetWorksheetOleObjectRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetOleObjects(data *DeleteWorksheetOleObjectsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetOleObject(data *DeleteWorksheetOleObjectRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWorksheetOleObject(data *PostUpdateWorksheetOleObjectRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetOleObject(data *PutWorksheetOleObjectRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetVerticalPageBreaks(data *GetVerticalPageBreaksRequest  ) (  VerticalPageBreaksResponse,  *http.Response, error) {
	var (
	localVarReturnValue VerticalPageBreaksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetHorizontalPageBreaks(data *GetHorizontalPageBreaksRequest  ) (  HorizontalPageBreaksResponse,  *http.Response, error) {
	var (
	localVarReturnValue HorizontalPageBreaksResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetVerticalPageBreak(data *GetVerticalPageBreakRequest  ) (  VerticalPageBreakResponse,  *http.Response, error) {
	var (
	localVarReturnValue VerticalPageBreakResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetHorizontalPageBreak(data *GetHorizontalPageBreakRequest  ) (  HorizontalPageBreakResponse,  *http.Response, error) {
	var (
	localVarReturnValue HorizontalPageBreakResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutVerticalPageBreak(data *PutVerticalPageBreakRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutHorizontalPageBreak(data *PutHorizontalPageBreakRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteVerticalPageBreaks(data *DeleteVerticalPageBreaksRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteHorizontalPageBreaks(data *DeleteHorizontalPageBreaksRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteVerticalPageBreak(data *DeleteVerticalPageBreakRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteHorizontalPageBreak(data *DeleteHorizontalPageBreakRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetPageSetup(data *GetPageSetupRequest  ) (  PageSetupResponse,  *http.Response, error) {
	var (
	localVarReturnValue PageSetupResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPageSetup(data *PostPageSetupRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteHeaderFooter(data *DeleteHeaderFooterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetHeader(data *GetHeaderRequest  ) (  PageSectionsResponse,  *http.Response, error) {
	var (
	localVarReturnValue PageSectionsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostHeader(data *PostHeaderRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetFooter(data *GetFooterRequest  ) (  PageSectionsResponse,  *http.Response, error) {
	var (
	localVarReturnValue PageSectionsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostFooter(data *PostFooterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostFitWideToPages(data *PostFitWideToPagesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostFitTallToPages(data *PostFitTallToPagesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPictures(data *GetWorksheetPicturesRequest  ) (  PicturesResponse,  *http.Response, error) {
	var (
	localVarReturnValue PicturesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPictureWithFormat(data *GetWorksheetPictureWithFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetAddPicture(data *PutWorksheetAddPictureRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) AddPictureInCell(data *AddPictureInCellRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetPicture(data *PostWorksheetPictureRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetPicture(data *DeleteWorksheetPictureRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetPictures(data *DeleteWorksheetPicturesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPivotTables(data *GetWorksheetPivotTablesRequest  ) (  PivotTablesResponse,  *http.Response, error) {
	var (
	localVarReturnValue PivotTablesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPivotTable(data *GetWorksheetPivotTableRequest  ) (  PivotTableResponse,  *http.Response, error) {
	var (
	localVarReturnValue PivotTableResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetPivotTableField(data *GetPivotTableFieldRequest  ) (  PivotFieldResponse,  *http.Response, error) {
	var (
	localVarReturnValue PivotFieldResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPivotTableFilters(data *GetWorksheetPivotTableFiltersRequest  ) (  PivotFiltersResponse,  *http.Response, error) {
	var (
	localVarReturnValue PivotFiltersResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPivotTableFilter(data *GetWorksheetPivotTableFilterRequest  ) (  PivotFilterResponse,  *http.Response, error) {
	var (
	localVarReturnValue PivotFilterResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetPivotTable(data *PutWorksheetPivotTableRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutPivotTableField(data *PutPivotTableFieldRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetPivotTableFilter(data *PutWorksheetPivotTableFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPivotTableFieldHideItem(data *PostPivotTableFieldHideItemRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPivotTableFieldMoveTo(data *PostPivotTableFieldMoveToRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPivotTableCellStyle(data *PostPivotTableCellStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPivotTableStyle(data *PostPivotTableStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPivotTableUpdatePivotFields(data *PostPivotTableUpdatePivotFieldsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostPivotTableUpdatePivotField(data *PostPivotTableUpdatePivotFieldRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetPivotTableCalculate(data *PostWorksheetPivotTableCalculateRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetPivotTableMove(data *PostWorksheetPivotTableMoveRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetPivotTables(data *DeleteWorksheetPivotTablesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetPivotTable(data *DeleteWorksheetPivotTableRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeletePivotTableField(data *DeletePivotTableFieldRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetPivotTableFilters(data *DeleteWorksheetPivotTableFiltersRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetPivotTableFilter(data *DeleteWorksheetPivotTableFilterRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetDocumentProperties(data *GetDocumentPropertiesRequest  ) (  CellsDocumentPropertiesResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsDocumentPropertiesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutDocumentProperty(data *PutDocumentPropertyRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetDocumentProperty(data *GetDocumentPropertyRequest  ) (  CellsDocumentPropertyResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsDocumentPropertyResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteDocumentProperty(data *DeleteDocumentPropertyRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteDocumentProperties(data *DeleteDocumentPropertiesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostDigitalSignature(data *PostDigitalSignatureRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostEncryptWorkbook(data *PostEncryptWorkbookRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteDecryptWorkbook(data *DeleteDecryptWorkbookRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostProtectWorkbook(data *PostProtectWorkbookRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteUnProtectWorkbook(data *DeleteUnProtectWorkbookRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutDocumentProtectFromChanges(data *PutDocumentProtectFromChangesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteDocumentUnProtectFromChanges(data *DeleteDocumentUnProtectFromChangesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUnlock(data *PostUnlockRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostLock(data *PostLockRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostProtect(data *PostProtectRequest  ) (  FilesResult,  *http.Response, error) {
	var (
	localVarReturnValue FilesResult 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangesCopy(data *PostWorksheetCellsRangesCopyRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeMerge(data *PostWorksheetCellsRangeMergeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeUnMerge(data *PostWorksheetCellsRangeUnMergeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeStyle(data *PostWorksheetCellsRangeStyleRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetCellsRangeValue(data *GetWorksheetCellsRangeValueRequest  ) (  RangeValueResponse,  *http.Response, error) {
	var (
	localVarReturnValue RangeValueResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeValue(data *PostWorksheetCellsRangeValueRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeMoveTo(data *PostWorksheetCellsRangeMoveToRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeSort(data *PostWorksheetCellsRangeSortRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeOutlineBorder(data *PostWorksheetCellsRangeOutlineBorderRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeColumnWidth(data *PostWorksheetCellsRangeColumnWidthRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeRowHeight(data *PostWorksheetCellsRangeRowHeightRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCellsRangeToImage(data *PostWorksheetCellsRangeToImageRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetCellsRange(data *PutWorksheetCellsRangeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetCellsRange(data *DeleteWorksheetCellsRangeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetShapes(data *GetWorksheetShapesRequest  ) (  ShapesResponse,  *http.Response, error) {
	var (
	localVarReturnValue ShapesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetShape(data *GetWorksheetShapeRequest  ) (  ShapeResponse,  *http.Response, error) {
	var (
	localVarReturnValue ShapeResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetShape(data *PutWorksheetShapeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetShapes(data *DeleteWorksheetShapesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetShape(data *DeleteWorksheetShapeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetShape(data *PostWorksheetShapeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetGroupShape(data *PostWorksheetGroupShapeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetUngroupShape(data *PostWorksheetUngroupShapeRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetSparklineGroups(data *GetWorksheetSparklineGroupsRequest  ) (  SparklineGroupsResponse,  *http.Response, error) {
	var (
	localVarReturnValue SparklineGroupsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetSparklineGroup(data *GetWorksheetSparklineGroupRequest  ) (  SparklineGroupResponse,  *http.Response, error) {
	var (
	localVarReturnValue SparklineGroupResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetSparklineGroups(data *DeleteWorksheetSparklineGroupsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetSparklineGroup(data *DeleteWorksheetSparklineGroupRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetSparklineGroup(data *PutWorksheetSparklineGroupRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetSparklineGroup(data *PostWorksheetSparklineGroupRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCharacterCount(data *PostCharacterCountRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWordsCount(data *PostWordsCountRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSpecifyWordsCount(data *PostSpecifyWordsCountRequest  ) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRunTask(data *PostRunTaskRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAddTextContent(data *PostAddTextContentRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostTrimContent(data *PostTrimContentRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWordCase(data *PostUpdateWordCaseRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRemoveCharacters(data *PostRemoveCharactersRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostConvertText(data *PostConvertTextRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRemoveDuplicates(data *PostRemoveDuplicatesRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostExtractText(data *PostExtractTextRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostSplitText(data *PostSplitTextRequest  ) (  FileInfo,  *http.Response, error) {
	var (
	localVarReturnValue FileInfo 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbookDefaultStyle(data *GetWorkbookDefaultStyleRequest  ) (  StyleResponse,  *http.Response, error) {
	var (
	localVarReturnValue StyleResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbookTextItems(data *GetWorkbookTextItemsRequest  ) (  TextItemsResponse,  *http.Response, error) {
	var (
	localVarReturnValue TextItemsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbookNames(data *GetWorkbookNamesRequest  ) (  NamesResponse,  *http.Response, error) {
	var (
	localVarReturnValue NamesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorkbookName(data *PutWorkbookNameRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbookName(data *GetWorkbookNameRequest  ) (  NameResponse,  *http.Response, error) {
	var (
	localVarReturnValue NameResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookName(data *PostWorkbookNameRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbookNameValue(data *GetWorkbookNameValueRequest  ) (  RangeValueResponse,  *http.Response, error) {
	var (
	localVarReturnValue RangeValueResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorkbookNames(data *DeleteWorkbookNamesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorkbookName(data *DeleteWorkbookNameRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbooksMerge(data *PostWorkbooksMergeRequest  ) (  WorkbookResponse,  *http.Response, error) {
	var (
	localVarReturnValue WorkbookResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbooksTextSearch(data *PostWorkbooksTextSearchRequest  ) (  TextItemsResponse,  *http.Response, error) {
	var (
	localVarReturnValue TextItemsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookTextReplace(data *PostWorkbookTextReplaceRequest  ) (  WorkbookReplaceResponse,  *http.Response, error) {
	var (
	localVarReturnValue WorkbookReplaceResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookGetSmartMarkerResult(data *PostWorkbookGetSmartMarkerResultRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorkbookCreate(data *PutWorkbookCreateRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookSplit(data *PostWorkbookSplitRequest  ) (  SplitResultResponse,  *http.Response, error) {
	var (
	localVarReturnValue SplitResultResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookCalculateFormula(data *PostWorkbookCalculateFormulaRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAutofitWorkbookRows(data *PostAutofitWorkbookRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAutofitWorkbookColumns(data *PostAutofitWorkbookColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorkbookSettings(data *GetWorkbookSettingsRequest  ) (  WorkbookSettingsResponse,  *http.Response, error) {
	var (
	localVarReturnValue WorkbookSettingsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorkbookSettings(data *PostWorkbookSettingsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorkbookBackground(data *PutWorkbookBackgroundRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorkbookBackground(data *DeleteWorkbookBackgroundRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorkbookWaterMarker(data *PutWorkbookWaterMarkerRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetPageCount(data *GetPageCountRequest  ) (  *int64,  *http.Response, error) {
	var (
	localVarReturnValue *int64 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetAllStyles(data *GetAllStylesRequest  ) (  StylesResponse,  *http.Response, error) {
	var (
	localVarReturnValue StylesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheets(data *GetWorksheetsRequest  ) (  WorksheetsResponse,  *http.Response, error) {
	var (
	localVarReturnValue WorksheetsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetWithFormat(data *GetWorksheetWithFormatRequest   , options ...CellsCloudOption) (  []byte,  *http.Response, error) {
	var (
	localVarReturnValue []byte 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
	 
		if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	if len(options) > 0 {
		for _, option := range options {
			if strings.EqualFold(option.OptionName, "LocalOutPath") {
				file, err := os.OpenFile(option.OptionValue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				if _, err := file.Write(localVarReturnValue); err != nil {
					panic(err)
				}
			}
		}
	}
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutChangeVisibilityWorksheet(data *PutChangeVisibilityWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutActiveWorksheet(data *PutActiveWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutInsertNewWorksheet(data *PutInsertNewWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutAddNewWorksheet(data *PutAddNewWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheet(data *DeleteWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheets(data *DeleteWorksheetsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostMoveWorksheet(data *PostMoveWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutProtectWorksheet(data *PutProtectWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteUnprotectWorksheet(data *DeleteUnprotectWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetTextItems(data *GetWorksheetTextItemsRequest  ) (  TextItemsResponse,  *http.Response, error) {
	var (
	localVarReturnValue TextItemsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetComments(data *GetWorksheetCommentsRequest  ) (  CommentsResponse,  *http.Response, error) {
	var (
	localVarReturnValue CommentsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetComment(data *GetWorksheetCommentRequest  ) (  CommentResponse,  *http.Response, error) {
	var (
	localVarReturnValue CommentResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetComment(data *PutWorksheetCommentRequest  ) (  CommentResponse,  *http.Response, error) {
	var (
	localVarReturnValue CommentResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetComment(data *PostWorksheetCommentRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetComment(data *DeleteWorksheetCommentRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetComments(data *DeleteWorksheetCommentsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetMergedCells(data *GetWorksheetMergedCellsRequest  ) (  MergedCellsResponse,  *http.Response, error) {
	var (
	localVarReturnValue MergedCellsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetMergedCell(data *GetWorksheetMergedCellRequest  ) (  MergedCellResponse,  *http.Response, error) {
	var (
	localVarReturnValue MergedCellResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetCalculateFormula(data *GetWorksheetCalculateFormulaRequest  ) (  SingleValueResponse,  *http.Response, error) {
	var (
	localVarReturnValue SingleValueResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetCalculateFormula(data *PostWorksheetCalculateFormulaRequest  ) (  SingleValueResponse,  *http.Response, error) {
	var (
	localVarReturnValue SingleValueResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetTextSearch(data *PostWorksheetTextSearchRequest  ) (  TextItemsResponse,  *http.Response, error) {
	var (
	localVarReturnValue TextItemsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetTextReplace(data *PostWorksheetTextReplaceRequest  ) (  WorksheetReplaceResponse,  *http.Response, error) {
	var (
	localVarReturnValue WorksheetReplaceResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetRangeSort(data *PostWorksheetRangeSortRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAutofitWorksheetRow(data *PostAutofitWorksheetRowRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAutofitWorksheetRows(data *PostAutofitWorksheetRowsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostAutofitWorksheetColumns(data *PostAutofitWorksheetColumnsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetBackground(data *PutWorksheetBackgroundRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetBackground(data *DeleteWorksheetBackgroundRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetFreezePanes(data *PutWorksheetFreezePanesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetFreezePanes(data *DeleteWorksheetFreezePanesRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostCopyWorksheet(data *PostCopyWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostRenameWorksheet(data *PostRenameWorksheetRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWorksheetProperty(data *PostUpdateWorksheetPropertyRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetNamedRanges(data *GetNamedRangesRequest  ) (  RangesResponse,  *http.Response, error) {
	var (
	localVarReturnValue RangesResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetNamedRangeValue(data *GetNamedRangeValueRequest  ) (  RangeValueResponse,  *http.Response, error) {
	var (
	localVarReturnValue RangeValueResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostUpdateWorksheetZoom(data *PostUpdateWorksheetZoomRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetPageCount(data *GetWorksheetPageCountRequest  ) (  *int64,  *http.Response, error) {
	var (
	localVarReturnValue *int64 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetValidations(data *GetWorksheetValidationsRequest  ) (  ValidationsResponse,  *http.Response, error) {
	var (
	localVarReturnValue ValidationsResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) GetWorksheetValidation(data *GetWorksheetValidationRequest  ) (  ValidationResponse,  *http.Response, error) {
	var (
	localVarReturnValue ValidationResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PutWorksheetValidation(data *PutWorksheetValidationRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) PostWorksheetValidation(data *PostWorksheetValidationRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetValidation(data *DeleteWorksheetValidationRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}


func (a *CellsApiService) DeleteWorksheetValidations(data *DeleteWorksheetValidationsRequest  ) (  CellsCloudResponse,  *http.Response, error) {
	var (
	localVarReturnValue CellsCloudResponse 

	)

    r, err := data.CreateRequestData(a.client);
    if err != nil {
        return  localVarReturnValue,  nil, err
    }

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue,  localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue,   localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}
		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
 
	return localVarReturnValue,  localVarHttpResponse, err
}

