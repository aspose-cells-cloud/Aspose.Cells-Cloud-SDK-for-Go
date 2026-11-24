/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="GradientFill">
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

// GradientFill            Represents the gradient fill.
type GradientFill struct {
	// Gets the gradient fill type.
	FillType string `json:"FillType,omitempty" xml:"FillType"`
	// Gets the gradient direction type.
	DirectionType string `json:"DirectionType,omitempty" xml:"DirectionType"`
	// The angle of linear fill.
	Angle *float64 `json:"Angle,omitempty" xml:"Angle"`
	// Represents the gradient stop collection.
	GradientStops []GradientFillStop `json:"GradientStops,omitempty" xml:"GradientStops"`
}
