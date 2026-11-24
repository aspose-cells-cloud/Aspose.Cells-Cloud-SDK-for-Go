/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DateTimeGroupItem">
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

// DateTimeGroupItem Represents the datetime's group setting.
type DateTimeGroupItem struct {
	// Gets and sets the group type.
	DateTimeGroupingType string `json:"DateTimeGroupingType,omitempty" xml:"DateTimeGroupingType"`
	// Gets and sets the day of the grouped date time.
	Day *int32 `json:"Day,omitempty" xml:"Day"`
	// Gets and sets the hour of the grouped date time.
	Hour *int32 `json:"Hour,omitempty" xml:"Hour"`
	// Gets and sets the minute of the grouped date time.
	Minute *int32 `json:"Minute,omitempty" xml:"Minute"`
	// Gets and sets the month of the grouped date time.
	Month *int32 `json:"Month,omitempty" xml:"Month"`
	// Gets and sets the second of the grouped date time.
	Second *int32 `json:"Second,omitempty" xml:"Second"`
	// Gets and sets the year of the grouped date time.
	Year *int32 `json:"Year,omitempty" xml:"Year"`
}
