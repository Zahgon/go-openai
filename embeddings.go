package openai

import (
	"context"
	"errors"
)

var ErrVectorLengthMismatch = errors.New("vector length mismatch")

type EmbeddingModel string

const (
	AdaSimilarity         EmbeddingModel = "text-similarity-ada-001"
	BabbageSimilarity     EmbeddingModel = "text-similarity-babbage-001"
	CurieSimilarity       EmbeddingModel = "text-similarity-curie-001"
	DavinciSimilarity     EmbeddingModel = "text-similarity-davinci-001"
	AdaSearchDocument     EmbeddingModel = "text-search-ada-doc-001"
	AdaSearchQuery        EmbeddingModel = "text-search-ada-query-001"
	BabbageSearchDocument EmbeddingModel = "text-search-babbage-doc-001"
	BabbageSearchQuery    EmbeddingModel = "text-search-babbage-query-001"
	CurieSearchDocument   EmbeddingModel = "text-search-curie-doc-001"
	CurieSearchQuery      EmbeddingModel = "text-search-curie-query-001"
	DavinciSearchDocument EmbeddingModel = "text-search-davinci-doc-001"
	DavinciSearchQuery    EmbeddingModel = "text-search-davinci-query-001"
	AdaCodeSearchCode     EmbeddingModel = "code-search-ada-code-001"
	AdaCodeSearchText     EmbeddingModel = "code-search-ada-text-001"
	BabbageCodeSearchCode EmbeddingModel = "code-search-babbage-code-001"
	BabbageCodeSearchText EmbeddingModel = "code-search-babbage-text-001"

	AdaEmbeddingV2  EmbeddingModel = "text-embedding-ada-002"
	SmallEmbedding3 EmbeddingModel = "text-embedding-3-small"
	LargeEmbedding3 EmbeddingModel = "text-embedding-3-large"
)

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

func (e *Embedding) DotProduct(other *Embedding) (float32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type EmbeddingResponse struct {
	Object string         `json:"object"`
	Data   []Embedding    `json:"data"`
	Model  EmbeddingModel `json:"model"`
	Usage  Usage          `json:"usage"`

	httpHeader
}

type base64String string

func (b base64String) Decode() ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

type Base64Embedding struct {
	Object    string       `json:"object"`
	Embedding base64String `json:"embedding"`
	Index     int          `json:"index"`
}

type EmbeddingResponseBase64 struct {
	Object string            `json:"object"`
	Data   []Base64Embedding `json:"data"`
	Model  EmbeddingModel    `json:"model"`
	Usage  Usage             `json:"usage"`

	httpHeader
}

func (r *EmbeddingResponseBase64) ToEmbeddingResponse() (EmbeddingResponse, error) {
	_ = "STUB: not implemented"
	return *new(EmbeddingResponse), nil
}

type EmbeddingRequestConverter interface {
	Convert() EmbeddingRequest
}

type EmbeddingEncodingFormat string

const (
	EmbeddingEncodingFormatFloat  EmbeddingEncodingFormat = "float"
	EmbeddingEncodingFormatBase64 EmbeddingEncodingFormat = "base64"
)

type EmbeddingRequest struct {
	Input          any                     `json:"input"`
	Model          EmbeddingModel          `json:"model"`
	User           string                  `json:"user,omitempty"`
	EncodingFormat EmbeddingEncodingFormat `json:"encoding_format,omitempty"`

	Dimensions int `json:"dimensions,omitempty"`

	ExtraBody map[string]any `json:"extra_body,omitempty"`
}

func (r EmbeddingRequest) Convert() EmbeddingRequest {
	_ = "STUB: not implemented"
	return *new(EmbeddingRequest)
}

type EmbeddingRequestStrings struct {
	Input []string `json:"input"`

	Model EmbeddingModel `json:"model"`

	User string `json:"user"`

	EncodingFormat EmbeddingEncodingFormat `json:"encoding_format,omitempty"`

	Dimensions int `json:"dimensions,omitempty"`

	ExtraBody map[string]any `json:"extra_body,omitempty"`
}

func (r EmbeddingRequestStrings) Convert() EmbeddingRequest {
	_ = "STUB: not implemented"
	return *new(EmbeddingRequest)
}

type EmbeddingRequestTokens struct {
	Input [][]int `json:"input"`

	Model EmbeddingModel `json:"model"`

	User string `json:"user"`

	EncodingFormat EmbeddingEncodingFormat `json:"encoding_format,omitempty"`

	Dimensions int `json:"dimensions,omitempty"`

	ExtraBody map[string]any `json:"extra_body,omitempty"`
}

func (r EmbeddingRequestTokens) Convert() EmbeddingRequest {
	_ = "STUB: not implemented"
	return *new(EmbeddingRequest)
}

func (c *Client) CreateEmbeddings(
	ctx context.Context,
	conv EmbeddingRequestConverter,
) (res EmbeddingResponse, err error) {
	_ = "STUB: not implemented"
	return *new(EmbeddingResponse), nil
}
