/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="configuration.go">
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
	"strings"
	"sync"
	"time"
)

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Version       string            `json:"version,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	ClientSecret  string            `json:"-"`
	ClientId      string
	AccessToken   string `json:"-"`

	GetAccessTokenTime time.Time
	TokenExpiresAt     time.Time

	tokenMu sync.RWMutex // protects AccessToken, GetAccessTokenTime, TokenExpiresAt
}

func NewConfiguration(clientId string, clientSecret string, basePath string, version string) *Configuration {
	cfg := &Configuration{
		// The SDK targets the v4.0 API; the legacy v3.0/v1.1 endpoints are opt-in
		// via the basePath suffix or the explicit version argument.
		BasePath:      "https://api.aspose.cloud/v4.0",
		Version:       "v4.0",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Aspose Cells Cloud SDK for Go",
		ClientSecret:  clientSecret,
		ClientId:      clientId,
	}
	if basePath != "" {
		cfg.BasePath = basePath
	}
	if strings.HasSuffix(cfg.BasePath, "/") {
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-1]
	}

	switch {
	case strings.HasSuffix(cfg.BasePath, "v4.0"):
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-5]
		cfg.Version = "v4.0"
	case strings.HasSuffix(cfg.BasePath, "v3.0"):
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-5]
		cfg.Version = "v3.0"
	case strings.HasSuffix(cfg.BasePath, "v1.1"):
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-5]
		cfg.Version = "v1.1"
	}

	if version != "" {
		cfg.Version = version
	}

	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
