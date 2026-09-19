package asposecellscloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Client SDK Core client
type AsposeCellsCloudClient struct {
	cfg        *Configuration
	httpClient *http.Client
	timeout    time.Duration
	retries    int
}

// Client configuration Option function type
type AsposeCellsCloudClientOption func(*AsposeCellsCloudClient)

// NewClient Create a new Client instance
func NewAsposeCellsCloudClient(clientId string, clientSecret string, baseURL string, opts ...AsposeCellsCloudClientOption) *AsposeCellsCloudClient {
	client := &AsposeCellsCloudClient{
		cfg:        NewConfiguration(clientId, clientSecret, baseURL, "v4.0"),
		httpClient: &http.Client{},
		timeout:    30 * time.Second,
		retries:    0,
	}

	// Apply all options
	for _, opt := range opts {
		opt(client)
	}

	// Set the default timeout. A caller-provided HTTP client (WithHTTPClient)
	// that already sets its own timeout keeps it.
	if client.httpClient.Timeout == 0 {
		client.httpClient.Timeout = client.timeout
	}

	return client
}

// WithHTTPClient overrides the default *http.Client used for all requests.
// Supply your own client to inject a custom Transport (mTLS, tracing, test
// mocks, ...). The configured request timeout is applied only when the given
// client does not set one itself.
func WithHTTPClient(hc *http.Client) AsposeCellsCloudClientOption {
	return func(c *AsposeCellsCloudClient) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithTimeout Set request timeout. The timeout is applied in the constructor
// only when the HTTP client does not already define its own (an injected
// client's non-zero Timeout takes precedence), which makes the effective
// timeout independent of option order.
func WithTimeout(timeout time.Duration) AsposeCellsCloudClientOption {
	return func(c *AsposeCellsCloudClient) {
		c.timeout = timeout
	}
}

// WithRetries Set the number of retries
func WithRetries(retries int) AsposeCellsCloudClientOption {
	return func(c *AsposeCellsCloudClient) {
		c.retries = retries
	}
}

// WithHeader Set global request headers
func WithHeader(key, value string) AsposeCellsCloudClientOption {
	return func(c *AsposeCellsCloudClient) {
		c.cfg.DefaultHeader[key] = value
	}
}

// tokenSafetyMargin is subtracted from the token lifetime so we refresh
// before the server-side expiration. 60 seconds is a reasonable buffer.
const tokenSafetyMargin = 60 * time.Second

// addAuth attaches a valid Bearer token to the request. It reuses the cached
// token when it has not expired yet and fetches a new one otherwise.
func (client *AsposeCellsCloudClient) addAuth(request *http.Request) error {
	client.cfg.tokenMu.RLock()
	token := client.cfg.AccessToken
	expiresAt := client.cfg.TokenExpiresAt
	client.cfg.tokenMu.RUnlock()

	// If the cached token is still valid (with safety margin), reuse it.
	if token != "" && time.Now().Before(expiresAt.Add(-tokenSafetyMargin)) {
		request.Header.Set("Authorization", "Bearer "+token)
		return nil
	}

	// Token is missing or about to expire — fetch a new one. The outgoing
	// request's context is reused so token refresh honors caller cancellation.
	if err := client.RequestOauthToken(request.Context()); err != nil {
		return fmt.Errorf("oauth token request failed: %w", err)
	}

	client.cfg.tokenMu.RLock()
	token = client.cfg.AccessToken
	client.cfg.tokenMu.RUnlock()

	request.Header.Set("Authorization", "Bearer "+token)
	return nil
}

// RequestOauthToken fetches a new OAuth2 access token from the Cells Cloud
// token endpoint and caches it together with its expiration time.
//
// The token request runs through the client's injected *http.Client, so a
// custom Transport (mTLS, proxy, tracing, test mocks), its Timeout, and the
// caller's context all apply to token refresh exactly as they do to the data
// plane requests.
func (client *AsposeCellsCloudClient) RequestOauthToken(ctx context.Context) error {
	// The OAuth token endpoint lives under the active API version prefix (v4.0
	// by default, see NewConfiguration). The v1.1 API uses the legacy oauth2
	// endpoint instead.
	getAccessTokenURI := client.cfg.BasePath + "/" + client.cfg.Version + "/cells/connect/token"

	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {client.cfg.ClientId},
		"client_secret": {client.cfg.ClientSecret},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, getAccessTokenURI, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("oauth token endpoint returned HTTP %d: %s", resp.StatusCode, string(body))
	}

	var tr TokenResponse
	if err = json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return err
	}

	now := time.Now()
	client.cfg.tokenMu.Lock()
	client.cfg.GetAccessTokenTime = now
	client.cfg.AccessToken = tr.AccessToken
	if tr.ExpiresIn > 0 {
		client.cfg.TokenExpiresAt = now.Add(time.Duration(tr.ExpiresIn) * time.Second)
	} else {
		// Default to 1 hour if the server does not report expiration.
		client.cfg.TokenExpiresAt = now.Add(1 * time.Hour)
	}
	client.cfg.tokenMu.Unlock()
	return nil
}

// Do Execute one or more requests (core scheduling method)
func (client *AsposeCellsCloudClient) Do(ctx context.Context, requests ...RequestOption) ([]*RichResponse, error) {
	var responses []*RichResponse

	for _, req := range requests {
		resp, err := client.executeWithRetry(ctx, req)
		if err != nil {
			return responses, &SDKError{
				Code:    -1,
				Message: "request execution failed",
				Err:     err,
			}
		}
		responses = append(responses, resp)
	}
	return responses, nil
}

// executeWithRetry Execution logic with retries.
//
// Missing required parameters are rejected here, before the retry loop, so
// deterministic validation errors fail fast instead of being retried. Only
// genuine execution failures (network, HTTP) trigger the backoff retries.
func (client *AsposeCellsCloudClient) executeWithRetry(ctx context.Context, req RequestOption) (*RichResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("%w: request is nil", ErrInvalidParam)
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var lastErr error
	for attempt := 0; attempt <= client.retries; attempt++ {
		resp, err := client.executeOnce(ctx, req)
		if err == nil {
			return resp, nil
		}
		lastErr = err
		if attempt < client.retries {
			// Simple exponential backoff (1s, 2s, 4s, ...), interrupted early when
			// the caller's context is done.
			if err := sleepWithContext(ctx, time.Duration(1<<attempt)*time.Second); err != nil {
				return nil, err
			}
		}
	}

	return nil, lastErr
}

// sleepWithContext sleeps for d, or returns ctx.Err() when the context is done
// first, so a canceled request is not held up by the backoff timer.
func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// executeOnce Single execution logic. Missing required parameters were already
// rejected by executeWithRetry before the retry loop, so no validation is
// repeated here.
func (client *AsposeCellsCloudClient) executeOnce(ctx context.Context, req RequestOption) (*RichResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("%w: request is nil", ErrInvalidParam)
	}

	// 1. Build URL
	// Generated request paths already carry the API version prefix of the
	// operation they call (e.g. "/v4.0/cells/convert/spreadsheet",
	// "/v3.0/cells/Book1.xlsx"), taken from the spec's per-operation APIVersion.
	// The client adds no version prefix of its own: an operation is always sent
	// to the version the service declares for it.
	u, err := url.Parse(client.cfg.BasePath + req.GetPath())
	if err != nil {
		return nil, err
	}
	// Concatenate Query Parameters
	query := u.Query()
	for k, v := range req.GetQueryParameters() {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	u.RawQuery = query.Encode()

	// 2. Build the request body
	jsonBody := req.GetJSONBody()
	multipartForm := req.GetMultipartForm()
	hasJSON := jsonBody != nil
	hasMultipart := multipartForm != nil && len(multipartForm) > 0
	var bodyReader io.Reader
	var contentType string

	if hasJSON && hasMultipart {
		// Serialize the JSON into a string and stick it into the form as a regular text field
		jsonBytes, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		// Agreement: Pass the entire JSON as a text field named 'request_body_data'
		multipartForm["request_body_data"] = string(jsonBytes)

		bodyReader, contentType, err = client.buildMultipartForm(multipartForm)
		if err != nil {
			return nil, err
		}
	} else if hasJSON {
		jsonBytes, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBytes)
		contentType = "application/json"
	} else if hasMultipart {
		bodyReader, contentType, err = client.buildMultipartForm(multipartForm)
		if err != nil {
			return nil, err
		}
	}

	// 3. Create an HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, req.GetMethod(), u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// 4. Set request headers
	httpReq.Header.Set("Content-Type", contentType)
	httpReq.Header.Set("User-Agent", client.cfg.UserAgent)
	for k, v := range client.cfg.DefaultHeader {
		httpReq.Header.Set(k, v)
	}
	httpReq.Header.Set("x-aspose-client", "go sdk")
	httpReq.Header.Set("x-aspose-client-version", globalCellsCloudSDKVersion)

	if err := client.addAuth(httpReq); err != nil {
		return nil, err
	}

	// 5. Execute request
	resp, err := client.httpClient.Do(httpReq)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	// 6.Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 7. Read and build RichResponse
	return &RichResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}

// buildMultipartForm Building a multipart/form-data request body
func (client *AsposeCellsCloudClient) buildMultipartForm(form map[string]interface{}) (io.Reader, string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for key, value := range form {
		switch v := value.(type) {
		case string:
			if strings.HasPrefix(key, "@") { // file
				err := addFile(writer, key[1:], v)
				if err != nil {
					return nil, "", err
				}
			} else { // form value
				// Plain text field
				if err := writer.WriteField(key, v); err != nil {
					return nil, "", err
				}
			}
		case []byte:
			// Binary file field. The generated requests expose only a field name
			// for byte payloads, so append ".bin" to the part filename: the v4.0
			// server keys off the multipart filename extension to detect the file
			// type, and a bare name (e.g. "datafile") makes import endpoints fail
			// with HTTP 500 "Object reference not set".
			part, err := writer.CreateFormFile(key, key+".bin")
			if err != nil {
				return nil, "", err
			}
			if _, err := part.Write(v); err != nil {
				return nil, "", err
			}
		default:
			// Unsupported multipart form value type; skip silently.
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return &buf, writer.FormDataContentType(), nil
}

func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
