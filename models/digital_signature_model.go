/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DigitalSignature">
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

// DigitalSignature Signature in file.
type DigitalSignature struct {
	// The purpose to signature.
	Comments string `json:"Comments,omitempty" xml:"Comments"`
	// The time when the document was signed.
	SignTime string `json:"SignTime,omitempty" xml:"SignTime"`
	// Specifies a GUID which can be cross-referenced with the GUID of the signature line stored in the document content. Default value is Empty (all zeroes) Guid.
	Id string `json:"Id,omitempty" xml:"Id"`
	// Specifies the text of actual signature in the digital signature. Default value is Empty.
	Password string `json:"Password,omitempty" xml:"Password"`
	// Specifies an image for the digital signature. Default value is null.
	Image []interface{} `json:"Image,omitempty" xml:"Image"`
	// Specifies the class ID of the signature provider. Default value is Empty (all zeroes) Guid.
	ProviderId string `json:"ProviderId,omitempty" xml:"ProviderId"`
	// If this digital signature is valid and the document has not been tampered with, this value will be true.
	IsValid *bool `json:"IsValid,omitempty" xml:"IsValid"`
	// XAdES type. Default value is None(XAdES is off).
	XAdESType string `json:"XAdESType,omitempty" xml:"XAdESType"`
}
