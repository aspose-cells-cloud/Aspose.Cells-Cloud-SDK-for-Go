/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="source.go">
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

package datasource

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// DataSource defines a data source interface.
//
// Ownership rule: when the data source opened the resource itself (for example
// the file opened by FilePathSource), ByteData closes it after reading; when
// the resource is provided by the caller (for example the stream in
// ReaderSource), the caller retains ownership and ByteData does not close it.
type DataSource interface {
	ByteData() ([]byte, error)
}

// DataSink defines a data sink interface.
type DataSink interface {
	Write() (io.WriteCloser, error)
}

// --- FilePathSource ---
type FilePathSource string

func (p FilePathSource) Open() (io.ReadCloser, error) {
	return os.Open(string(p))
}

func (p FilePathSource) ByteData() ([]byte, error) {
	reader, err := p.Open()
	if err != nil {
		return nil, err
	}
	// This data source opened the file itself, so it closes it after reading.
	defer reader.Close()
	return io.ReadAll(reader)
}

// --- BytesSource ---
type BytesSource []byte

func (b BytesSource) Open() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(b)), nil
}

func (b BytesSource) ByteData() ([]byte, error) {
	return b, nil
}

// --- ReaderSource ---
type ReaderSource struct {
	reader io.ReadCloser
}

// NewReaderSource creates a ReaderSource from an io.ReadCloser.
func NewReaderSource(r io.ReadCloser) *ReaderSource {
	return &ReaderSource{reader: r}
}

func (r ReaderSource) Open() (io.ReadCloser, error) {
	return r.reader, nil
}

func (r ReaderSource) ByteData() ([]byte, error) {
	// The stream is provided by the caller, which retains ownership; read it
	// without closing it.
	return io.ReadAll(r.reader)
}

// --- FilePathSink ---
type FilePathSink string

func (p FilePathSink) Write() (io.WriteCloser, error) {
	return os.Create(string(p))
}

// --- BytesSink ---
type BytesSink struct {
	buf bytes.Buffer
}

func (b *BytesSink) Write() (io.WriteCloser, error) {
	// Reset() was removed to avoid a data race.
	// To reuse a BytesSink, create a new one instead.
	return nopWriteCloser{&b.buf}, nil
}

func (b *BytesSink) Bytes() []byte {
	return b.buf.Bytes()
}

// nopWriteCloser wraps an io.Writer.
type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error { return nil }

// UrlPathSink is an implementation of a data sink based on HTTP/HTTPS URLs.
// It implements the DataSink interface and is used to upload data to a remote server
// via an HTTP PUT request. This is useful for pushing converted files directly to
// your own backend, or to pre-signed cloud storage URLs (e.g., AWS S3, Aliyun OSS).
type UrlPathSink string

// Write initiates an HTTP PUT request to the specified URL and returns a WriteCloser.
// The caller should write data to this WriteCloser and then call Close().
// Upon Close(), the buffered data is flushed to the remote server.
func (u UrlPathSink) Write() (io.WriteCloser, error) {
	pr, pw := io.Pipe()

	// Start a background goroutine that streams data written to the pipe into
	// the HTTP request.
	go func() {
		defer pr.Close()

		req, err := http.NewRequest(http.MethodPut, string(u), pr)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		// Set the common streaming upload header.
		req.Header.Set("Content-Type", "application/octet-stream")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		defer resp.Body.Close()

		// Check the HTTP status code; treat any non-2xx response as an upload
		// failure.
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			pw.CloseWithError(fmt.Errorf("upload to URL failed with status %d", resp.StatusCode))
			return
		}
	}()

	return pw, nil
}
