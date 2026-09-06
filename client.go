package openai

import (
	"context"
	"io"
	"net/http"

	utils "github.com/sashabaranov/go-openai/internal"
)

type Client struct {
	config ClientConfig

	requestBuilder    utils.RequestBuilder
	createFormBuilder func(io.Writer) utils.FormBuilder
}

type Response interface {
	SetHeader(http.Header)
}

type httpHeader http.Header

func (h *httpHeader) SetHeader(header http.Header) { _ = "STUB: not implemented"; return }

func (h *httpHeader) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (h *httpHeader) GetRateLimitHeaders() RateLimitHeaders {
	_ = "STUB: not implemented"
	return *new(RateLimitHeaders)
}

type RawResponse struct {
	io.ReadCloser

	httpHeader
}

func NewClient(authToken string) *Client { _ = "STUB: not implemented"; return nil }

func NewClientWithConfig(config ClientConfig) *Client { _ = "STUB: not implemented"; return nil }

func NewOrgClient(authToken, org string) *Client { _ = "STUB: not implemented"; return nil }

type requestOptions struct {
	body   any
	header http.Header
}

type requestOption func(*requestOptions)

func withBody(body any) requestOption { _ = "STUB: not implemented"; return *new(requestOption) }

func withExtraBody(extraBody map[string]any) requestOption {
	_ = "STUB: not implemented"
	return *new(requestOption)
}

func withContentType(contentType string) requestOption {
	_ = "STUB: not implemented"
	return *new(requestOption)
}

func withBetaAssistantVersion(version string) requestOption {
	_ = "STUB: not implemented"
	return *new(requestOption)
}

func (c *Client) newRequest(ctx context.Context, method, url string, setters ...requestOption) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) sendRequest(req *http.Request, v Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) sendRequestRaw(req *http.Request) (response RawResponse, err error) {
	_ = "STUB: not implemented"
	return *new(RawResponse), nil
}

//nolint:bodyclose // body should be closed by outer function

func sendRequestStream[T streamable](client *Client, req *http.Request) (*streamReader[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:bodyclose // body is closed in stream.Close()

func (c *Client) setCommonHeaders(req *http.Request) { _ = "STUB: not implemented"; return }

func isFailureStatusCode(resp *http.Response) bool { _ = "STUB: not implemented"; return false }

func decodeResponse(body io.Reader, v any) error { _ = "STUB: not implemented"; return nil }

func decodeString(body io.Reader, output *string) error { _ = "STUB: not implemented"; return nil }

type fullURLOptions struct {
	model string
}

type fullURLOption func(*fullURLOptions)

func withModel(model string) fullURLOption { _ = "STUB: not implemented"; return *new(fullURLOption) }

var azureDeploymentsEndpoints = []string{
	completionsSuffix,
	"/embeddings",
	chatCompletionsSuffix,
	"/audio/transcriptions",
	"/audio/translations",
	"/audio/speech",
	"/images/generations",
}

func (c *Client) fullURL(suffix string, setters ...fullURLOption) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Client) suffixWithAPIVersion(suffix string) string { _ = "STUB: not implemented"; return "" }

func (c *Client) baseURLWithAzureDeployment(baseURL, suffix, model string) (newBaseURL string) {
	_ = "STUB: not implemented"
	return ""
}

func (c *Client) handleErrorResp(resp *http.Response) error { _ = "STUB: not implemented"; return nil }

func containsSubstr(s []string, e string) bool { _ = "STUB: not implemented"; return false }
