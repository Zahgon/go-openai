package openai

import (
	"context"
)

type FileRequest struct {
	FileName string `json:"file"`
	FilePath string `json:"-"`
	Purpose  string `json:"purpose"`
}

type PurposeType string

const (
	PurposeFineTune         PurposeType = "fine-tune"
	PurposeFineTuneResults  PurposeType = "fine-tune-results"
	PurposeAssistants       PurposeType = "assistants"
	PurposeAssistantsOutput PurposeType = "assistants_output"
	PurposeBatch            PurposeType = "batch"
)

type FileBytesRequest struct {
	Name string

	Bytes []byte

	Purpose PurposeType
}

type File struct {
	Bytes         int    `json:"bytes"`
	CreatedAt     int64  `json:"created_at"`
	ID            string `json:"id"`
	FileName      string `json:"filename"`
	Object        string `json:"object"`
	Status        string `json:"status"`
	Purpose       string `json:"purpose"`
	StatusDetails string `json:"status_details"`

	httpHeader
}

type FilesList struct {
	Files []File `json:"data"`

	httpHeader
}

func (c *Client) CreateFileBytes(ctx context.Context, request FileBytesRequest) (file File, err error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

func (c *Client) CreateFile(ctx context.Context, request FileRequest) (file File, err error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

func (c *Client) DeleteFile(ctx context.Context, fileID string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ListFiles(ctx context.Context) (files FilesList, err error) {
	_ = "STUB: not implemented"
	return *new(FilesList), nil
}

func (c *Client) GetFile(ctx context.Context, fileID string) (file File, err error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

func (c *Client) GetFileContent(ctx context.Context, fileID string) (content RawResponse, err error) {
	_ = "STUB: not implemented"
	return *new(RawResponse), nil
}
