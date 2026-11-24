/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="request_config.go">
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

package requests

type requestConfig struct {
	Params           map[string]interface{}
	extraQueryParams map[string]string
}

type Option interface {
	apply(*requestConfig)
}

type optionFunc func(*requestConfig)

func (f optionFunc) apply(c *requestConfig) {
	f(c)
}
func WithCommonParameter(key string, value interface{}) Option {

	return optionFunc(func(c *requestConfig) {
		if c.Params == nil {
			c.Params = make(map[string]interface{})
		}
		c.Params[key] = value
	})
}

// WithQueryParameter adds a single custom query parameter (name/value pair) to the
// request. This is an implicit extension on top of the operation's declared query
// parameters, used to pass extra attributes such as save options.
func WithQueryParameter(name, value string) Option {
	return optionFunc(func(c *requestConfig) {
		if c.extraQueryParams == nil {
			c.extraQueryParams = make(map[string]string)
		}
		c.extraQueryParams[name] = value
	})
}

// WithQueryParameters adds multiple custom query parameters at once.
func WithQueryParameters(params map[string]string) Option {
	return optionFunc(func(c *requestConfig) {
		if c.extraQueryParams == nil {
			c.extraQueryParams = make(map[string]string)
		}
		for k, v := range params {
			c.extraQueryParams[k] = v
		}
	})
}
