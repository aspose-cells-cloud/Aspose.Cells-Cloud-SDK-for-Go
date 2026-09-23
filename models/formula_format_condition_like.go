/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="formula_format_condition_like.go">
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

// FormulaFormatConditionLike accepts FormulaFormatCondition itself or any model that embeds
// it -- its only subclass is TextFormatCondition. Go has no subclass relation between
// structs, so a child cannot be assigned to a FormulaFormatCondition field or parameter;
// this interface is the spelling for "FormulaFormatCondition or any of its subclasses".
// Embedding promotes the marker method, so every subclass already satisfies it: hand it the
// concrete one, carrying the data the service routes on.
type FormulaFormatConditionLike interface {
	formulaFormatConditionMarker()
}
