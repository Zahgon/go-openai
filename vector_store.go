package openai

import (
	"context"
)

const (
	vectorStoresSuffix            = "/vector_stores"
	vectorStoresFilesSuffix       = "/files"
	vectorStoresFileBatchesSuffix = "/file_batches"
)

type VectorStoreFileCount struct {
	InProgress int `json:"in_progress"`
	Completed  int `json:"completed"`
	Failed     int `json:"failed"`
	Cancelled  int `json:"cancelled"`
	Total      int `json:"total"`
}

type VectorStore struct {
	ID           string               `json:"id"`
	Object       string               `json:"object"`
	CreatedAt    int64                `json:"created_at"`
	Name         string               `json:"name"`
	UsageBytes   int                  `json:"usage_bytes"`
	FileCounts   VectorStoreFileCount `json:"file_counts"`
	Status       string               `json:"status"`
	ExpiresAfter *VectorStoreExpires  `json:"expires_after"`
	ExpiresAt    *int                 `json:"expires_at"`
	Metadata     map[string]any       `json:"metadata"`

	httpHeader
}

type VectorStoreExpires struct {
	Anchor string `json:"anchor"`
	Days   int    `json:"days"`
}

type VectorStoreRequest struct {
	Name         string              `json:"name,omitempty"`
	FileIDs      []string            `json:"file_ids,omitempty"`
	ExpiresAfter *VectorStoreExpires `json:"expires_after,omitempty"`
	Metadata     map[string]any      `json:"metadata,omitempty"`
}

type VectorStoresList struct {
	VectorStores []VectorStore `json:"data"`
	LastID       *string       `json:"last_id"`
	FirstID      *string       `json:"first_id"`
	HasMore      bool          `json:"has_more"`
	httpHeader
}

type VectorStoreDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`

	httpHeader
}

type VectorStoreFile struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	CreatedAt     int64  `json:"created_at"`
	VectorStoreID string `json:"vector_store_id"`
	UsageBytes    int    `json:"usage_bytes"`
	Status        string `json:"status"`

	httpHeader
}

type VectorStoreFileRequest struct {
	FileID string `json:"file_id"`
}

type VectorStoreFilesList struct {
	VectorStoreFiles []VectorStoreFile `json:"data"`
	FirstID          *string           `json:"first_id"`
	LastID           *string           `json:"last_id"`
	HasMore          bool              `json:"has_more"`

	httpHeader
}

type VectorStoreFileBatch struct {
	ID            string               `json:"id"`
	Object        string               `json:"object"`
	CreatedAt     int64                `json:"created_at"`
	VectorStoreID string               `json:"vector_store_id"`
	Status        string               `json:"status"`
	FileCounts    VectorStoreFileCount `json:"file_counts"`

	httpHeader
}

type VectorStoreFileBatchRequest struct {
	FileIDs []string `json:"file_ids"`
}

func (c *Client) CreateVectorStore(ctx context.Context, request VectorStoreRequest) (response VectorStore, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStore), nil
}

func (c *Client) RetrieveVectorStore(
	ctx context.Context,
	vectorStoreID string,
) (response VectorStore, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStore), nil
}

func (c *Client) ModifyVectorStore(
	ctx context.Context,
	vectorStoreID string,
	request VectorStoreRequest,
) (response VectorStore, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStore), nil
}

func (c *Client) DeleteVectorStore(
	ctx context.Context,
	vectorStoreID string,
) (response VectorStoreDeleteResponse, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreDeleteResponse), nil
}

func (c *Client) ListVectorStores(
	ctx context.Context,
	pagination Pagination,
) (response VectorStoresList, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoresList), nil
}

func (c *Client) CreateVectorStoreFile(
	ctx context.Context,
	vectorStoreID string,
	request VectorStoreFileRequest,
) (response VectorStoreFile, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFile), nil
}

func (c *Client) RetrieveVectorStoreFile(
	ctx context.Context,
	vectorStoreID string,
	fileID string,
) (response VectorStoreFile, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFile), nil
}

func (c *Client) DeleteVectorStoreFile(
	ctx context.Context,
	vectorStoreID string,
	fileID string,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ListVectorStoreFiles(
	ctx context.Context,
	vectorStoreID string,
	pagination Pagination,
) (response VectorStoreFilesList, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFilesList), nil
}

func (c *Client) CreateVectorStoreFileBatch(
	ctx context.Context,
	vectorStoreID string,
	request VectorStoreFileBatchRequest,
) (response VectorStoreFileBatch, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFileBatch), nil
}

func (c *Client) RetrieveVectorStoreFileBatch(
	ctx context.Context,
	vectorStoreID string,
	batchID string,
) (response VectorStoreFileBatch, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFileBatch), nil
}

func (c *Client) CancelVectorStoreFileBatch(
	ctx context.Context,
	vectorStoreID string,
	batchID string,
) (response VectorStoreFileBatch, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFileBatch), nil
}

func (c *Client) ListVectorStoreFilesInBatch(
	ctx context.Context,
	vectorStoreID string,
	batchID string,
	pagination Pagination,
) (response VectorStoreFilesList, err error) {
	_ = "STUB: not implemented"
	return *new(VectorStoreFilesList), nil
}
