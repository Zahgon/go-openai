package openai

import (
	"context"
)

const batchesSuffix = "/batches"

type BatchEndpoint string

const (
	BatchEndpointChatCompletions BatchEndpoint = "/v1/chat/completions"
	BatchEndpointCompletions     BatchEndpoint = "/v1/completions"
	BatchEndpointEmbeddings      BatchEndpoint = "/v1/embeddings"
)

type BatchLineItem interface {
	MarshalBatchLineItem() []byte
}

type BatchChatCompletionRequest struct {
	CustomID string                `json:"custom_id"`
	Body     ChatCompletionRequest `json:"body"`
	Method   string                `json:"method"`
	URL      BatchEndpoint         `json:"url"`
}

func (r BatchChatCompletionRequest) MarshalBatchLineItem() []byte {
	_ = "STUB: not implemented"
	return nil
}

type BatchCompletionRequest struct {
	CustomID string            `json:"custom_id"`
	Body     CompletionRequest `json:"body"`
	Method   string            `json:"method"`
	URL      BatchEndpoint     `json:"url"`
}

func (r BatchCompletionRequest) MarshalBatchLineItem() []byte {
	_ = "STUB: not implemented"
	return nil
}

type BatchEmbeddingRequest struct {
	CustomID string           `json:"custom_id"`
	Body     EmbeddingRequest `json:"body"`
	Method   string           `json:"method"`
	URL      BatchEndpoint    `json:"url"`
}

func (r BatchEmbeddingRequest) MarshalBatchLineItem() []byte { _ = "STUB: not implemented"; return nil }

type Batch struct {
	ID       string        `json:"id"`
	Object   string        `json:"object"`
	Endpoint BatchEndpoint `json:"endpoint"`
	Errors   *struct {
		Object string `json:"object,omitempty"`
		Data   []struct {
			Code    string  `json:"code,omitempty"`
			Message string  `json:"message,omitempty"`
			Param   *string `json:"param,omitempty"`
			Line    *int    `json:"line,omitempty"`
		} `json:"data"`
	} `json:"errors"`
	InputFileID      string             `json:"input_file_id"`
	CompletionWindow string             `json:"completion_window"`
	Status           string             `json:"status"`
	OutputFileID     *string            `json:"output_file_id"`
	ErrorFileID      *string            `json:"error_file_id"`
	CreatedAt        int                `json:"created_at"`
	InProgressAt     *int               `json:"in_progress_at"`
	ExpiresAt        *int               `json:"expires_at"`
	FinalizingAt     *int               `json:"finalizing_at"`
	CompletedAt      *int               `json:"completed_at"`
	FailedAt         *int               `json:"failed_at"`
	ExpiredAt        *int               `json:"expired_at"`
	CancellingAt     *int               `json:"cancelling_at"`
	CancelledAt      *int               `json:"cancelled_at"`
	RequestCounts    BatchRequestCounts `json:"request_counts"`
	Metadata         map[string]any     `json:"metadata"`
}

type BatchRequestCounts struct {
	Total     int `json:"total"`
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
}

type CreateBatchRequest struct {
	InputFileID      string         `json:"input_file_id"`
	Endpoint         BatchEndpoint  `json:"endpoint"`
	CompletionWindow string         `json:"completion_window"`
	Metadata         map[string]any `json:"metadata"`
}

type BatchResponse struct {
	httpHeader
	Batch
}

func (c *Client) CreateBatch(
	ctx context.Context,
	request CreateBatchRequest,
) (response BatchResponse, err error) {
	_ = "STUB: not implemented"
	return *new(BatchResponse), nil
}

type UploadBatchFileRequest struct {
	FileName string
	Lines    []BatchLineItem
}

func (r *UploadBatchFileRequest) MarshalJSONL() []byte { _ = "STUB: not implemented"; return nil }

func (r *UploadBatchFileRequest) AddChatCompletion(customerID string, body ChatCompletionRequest) {
	_ = "STUB: not implemented"
	return
}

func (r *UploadBatchFileRequest) AddCompletion(customerID string, body CompletionRequest) {
	_ = "STUB: not implemented"
	return
}

func (r *UploadBatchFileRequest) AddEmbedding(customerID string, body EmbeddingRequest) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) UploadBatchFile(ctx context.Context, request UploadBatchFileRequest) (File, error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

type CreateBatchWithUploadFileRequest struct {
	Endpoint         BatchEndpoint  `json:"endpoint"`
	CompletionWindow string         `json:"completion_window"`
	Metadata         map[string]any `json:"metadata"`
	UploadBatchFileRequest
}

func (c *Client) CreateBatchWithUploadFile(
	ctx context.Context,
	request CreateBatchWithUploadFileRequest,
) (response BatchResponse, err error) {
	_ = "STUB: not implemented"
	return *new(BatchResponse), nil
}

func (c *Client) RetrieveBatch(
	ctx context.Context,
	batchID string,
) (response BatchResponse, err error) {
	_ = "STUB: not implemented"
	return *new(BatchResponse), nil
}

func (c *Client) CancelBatch(
	ctx context.Context,
	batchID string,
) (response BatchResponse, err error) {
	_ = "STUB: not implemented"
	return *new(BatchResponse), nil
}

type ListBatchResponse struct {
	httpHeader
	Object  string  `json:"object"`
	Data    []Batch `json:"data"`
	FirstID string  `json:"first_id"`
	LastID  string  `json:"last_id"`
	HasMore bool    `json:"has_more"`
}

func (c *Client) ListBatch(ctx context.Context, after *string, limit *int) (response ListBatchResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ListBatchResponse), nil
}
