package openai

import (
	"bufio"
	"net/http"
	"regexp"

	utils "github.com/sashabaranov/go-openai/internal"
)

var (
	headerData  = regexp.MustCompile(`^data:\s*`)
	errorPrefix = regexp.MustCompile(`^data:\s*{"error":`)
)

type streamable interface {
	ChatCompletionStreamResponse | CompletionResponse
}

type streamReader[T streamable] struct {
	emptyMessagesLimit uint
	isFinished         bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator utils.ErrorAccumulator
	unmarshaler    utils.Unmarshaler

	httpHeader
}

func (stream *streamReader[T]) Recv() (response T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (stream *streamReader[T]) RecvRaw() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gocognit
func (stream *streamReader[T]) processLines() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (stream *streamReader[T]) unmarshalError() (errResp *ErrorResponse) {
	_ = "STUB: not implemented"
	return nil
}

func (stream *streamReader[T]) Close() error { _ = "STUB: not implemented"; return nil }
