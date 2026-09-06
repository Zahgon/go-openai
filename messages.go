package openai

import (
	"context"
)

const (
	messagesSuffix = "messages"
)

type Message struct {
	ID          string           `json:"id"`
	Object      string           `json:"object"`
	CreatedAt   int              `json:"created_at"`
	ThreadID    string           `json:"thread_id"`
	Role        string           `json:"role"`
	Content     []MessageContent `json:"content"`
	FileIds     []string         `json:"file_ids"` //nolint:revive //backwards-compatibility
	AssistantID *string          `json:"assistant_id,omitempty"`
	RunID       *string          `json:"run_id,omitempty"`
	Metadata    map[string]any   `json:"metadata"`

	httpHeader
}

type MessagesList struct {
	Messages []Message `json:"data"`

	Object  string  `json:"object"`
	FirstID *string `json:"first_id"`
	LastID  *string `json:"last_id"`
	HasMore bool    `json:"has_more"`

	httpHeader
}

type MessageContent struct {
	Type      string       `json:"type"`
	Text      *MessageText `json:"text,omitempty"`
	ImageFile *ImageFile   `json:"image_file,omitempty"`
	ImageURL  *ImageURL    `json:"image_url,omitempty"`
}
type MessageText struct {
	Value       string `json:"value"`
	Annotations []any  `json:"annotations"`
}

type ImageFile struct {
	FileID string `json:"file_id"`
}

type ImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail"`
}

type MessageRequest struct {
	Role        string             `json:"role"`
	Content     string             `json:"content"`
	FileIds     []string           `json:"file_ids,omitempty"` //nolint:revive // backwards-compatibility
	Metadata    map[string]any     `json:"metadata,omitempty"`
	Attachments []ThreadAttachment `json:"attachments,omitempty"`
}

type MessageFile struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int    `json:"created_at"`
	MessageID string `json:"message_id"`

	httpHeader
}

type MessageFilesList struct {
	MessageFiles []MessageFile `json:"data"`

	httpHeader
}

type MessageDeletionStatus struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`

	httpHeader
}

func (c *Client) CreateMessage(ctx context.Context, threadID string, request MessageRequest) (msg Message, err error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func (c *Client) ListMessage(ctx context.Context, threadID string,
	limit *int,
	order *string,
	after *string,
	before *string,
	runID *string,
) (messages MessagesList, err error) {
	_ = "STUB: not implemented"
	return *new(MessagesList), nil
}

func (c *Client) RetrieveMessage(
	ctx context.Context,
	threadID, messageID string,
) (msg Message, err error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func (c *Client) ModifyMessage(
	ctx context.Context,
	threadID, messageID string,
	metadata map[string]string,
) (msg Message, err error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func (c *Client) RetrieveMessageFile(
	ctx context.Context,
	threadID, messageID, fileID string,
) (file MessageFile, err error) {
	_ = "STUB: not implemented"
	return *new(MessageFile), nil
}

func (c *Client) ListMessageFiles(
	ctx context.Context,
	threadID, messageID string,
) (files MessageFilesList, err error) {
	_ = "STUB: not implemented"
	return *new(MessageFilesList), nil
}

func (c *Client) DeleteMessage(
	ctx context.Context,
	threadID, messageID string,
) (status MessageDeletionStatus, err error) {
	_ = "STUB: not implemented"
	return *new(MessageDeletionStatus), nil
}
