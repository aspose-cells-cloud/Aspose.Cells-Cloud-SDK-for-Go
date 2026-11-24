/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ChartOperateParameter">
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

package models

import ()

// ChartOperateParameter Represents chart operate parameter.
type ChartOperateParameter struct {
	OperateParameter
	// Represents chart index.
	ChartIndex *int32 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
	// Represents chart type.
	ChartType string `json:"ChartType,omitempty" xml:"ChartType"`
	// Represents upper left row index of chart.
	UpperLeftRow *int32 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
	// Represents upper left column index of chart.
	UpperLeftColumn *int32 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
	// Represents lower right row index of chart.
	LowerRightRow *int32 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
	// Represents lower right column index of chart.
	LowerRightColumn *int32 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
	// Represents chart area.
	Area string `json:"Area,omitempty" xml:"Area"`
	// Represents whether to plot the series from a range of cell values by row or by column.
	IsVertical *bool `json:"IsVertical,omitempty" xml:"IsVertical"`
	// Represents chart category data.
	CategoryData string `json:"CategoryData,omitempty" xml:"CategoryData"`
	// Represents whether auto get serial name.
	IsAutoGetSerialName *bool `json:"IsAutoGetSerialName,omitempty" xml:"IsAutoGetSerialName"`
	// Represents chart title.
	Title string `json:"Title,omitempty" xml:"Title"`
}
