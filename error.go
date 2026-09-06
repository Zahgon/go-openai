package openai

type APIError struct {
	Code           any         `json:"code,omitempty"`
	Message        string      `json:"message"`
	Param          *string     `json:"param,omitempty"`
	Type           string      `json:"type"`
	HTTPStatus     string      `json:"-"`
	HTTPStatusCode int         `json:"-"`
	InnerError     *InnerError `json:"innererror,omitempty"`
}

type InnerError struct {
	Code                 string               `json:"code,omitempty"`
	ContentFilterResults ContentFilterResults `json:"content_filter_result,omitempty"`
}

type RequestError struct {
	HTTPStatus     string
	HTTPStatusCode int
	Err            error
	Body           []byte
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *APIError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *APIError) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (e *RequestError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *RequestError) Unwrap() error { _ = "STUB: not implemented"; return nil }
