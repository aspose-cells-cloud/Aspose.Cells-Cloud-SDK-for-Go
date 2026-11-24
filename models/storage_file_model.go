/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="StorageFile">
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

import (
	"time"
)

// StorageFile File or folder information
type StorageFile struct {
	// File or folder name.
	Name string `json:"Name,omitempty" xml:"Name"`
	// True if it is a folder.
	IsFolder *bool `json:"IsFolder,omitempty" xml:"IsFolder"`
	// File or folder last modified DateTime.
	ModifiedDate time.Time `json:"ModifiedDate,omitempty" xml:"ModifiedDate"`
	// File or folder size.
	Size *int64 `json:"Size,omitempty" xml:"Size"`
	// File or folder path.
	Path string `json:"Path,omitempty" xml:"Path"`
}

// storageFileMarker closes the StorageFileLike family: only StorageFile and the models
// embedding it carry this promoted method, so no unrelated type satisfies StorageFileLike.
func (m StorageFile) storageFileMarker() {}

// StorageFileLike accepts StorageFile itself or any model that embeds it -- its subclasses
// are FileVersion and GoogleDriveStorageFile. Go has no subclass relation between structs,
// so a child cannot be assigned to a StorageFile field or parameter; this interface is the
// spelling for "StorageFile or any of its subclasses". Embedding promotes the marker
// method, so every subclass already satisfies it: hand it the concrete one, carrying the
// data the service routes on.
type StorageFileLike interface {
	storageFileMarker()
}
