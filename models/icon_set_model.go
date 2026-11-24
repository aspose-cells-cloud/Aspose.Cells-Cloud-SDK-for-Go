/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="IconSet">
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

// IconSet Describe the IconSet conditional formatting rule. This conditional formatting     rule applies icons to cells according to their values.
type IconSet struct {
	// Get the from the collection
	CfIcons []ConditionalFormattingIcon `json:"CfIcons,omitempty" xml:"CfIcons"`
	// Get the CFValueObjects instance.
	Cfvos []ConditionalFormattingValue `json:"Cfvos,omitempty" xml:"Cfvos"`
	// Indicates whether the icon set is custom.             Default value is false.
	IsCustom *bool `json:"IsCustom,omitempty" xml:"IsCustom"`
	// Get or set the flag indicating whether to reverses the default order of the icons in this icon set.             Default value is false.
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse"`
	// Get or set the flag indicating whether to show the values of the cells on which this icon set is applied.             Default value is true.
	ShowValue *bool `json:"ShowValue,omitempty" xml:"ShowValue"`
	// Get or Set the icon set type to display.  Setting the type will auto check    if the current Cfvos's count is accord with the new type. If not accord,    old Cfvos will be cleaned and default Cfvos will be added.
	IconSetType string `json:"IconSetType,omitempty" xml:"IconSetType"`
}
