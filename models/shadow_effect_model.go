/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ShadowEffect">
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

// ShadowEffect            This class specifies the shadow effect of the chart element or shape.
type ShadowEffect struct {
	// Gets and sets the lighting angle. Range from 0 to 359.9 degrees.
	Angle *float64 `json:"Angle,omitempty" xml:"Angle"`
	// Gets and sets the blur of the shadow. Range from 0 to 100 points.
	Blur *float64 `json:"Blur,omitempty" xml:"Blur"`
	// Gets and sets the color of the shadow.
	Color *CellsColor `json:"Color,omitempty" xml:"Color"`
	// Gets and sets the distance of the shadow. Range from 0 to 200 points.
	Distance *float64 `json:"Distance,omitempty" xml:"Distance"`
	// Gets and sets the preset shadow type of the shadow.
	PresetType string `json:"PresetType,omitempty" xml:"PresetType"`
	// Gets and sets the size of the shadow. Range from 0 to 2.0.              Meaningless in inner shadow.
	Size *float64 `json:"Size,omitempty" xml:"Size"`
	// Gets and sets the degree of transparency of the shadow. Range from 0.0 (opaque) to 1.0 (clear).
	Transparency *float64 `json:"Transparency,omitempty" xml:"Transparency"`
}
