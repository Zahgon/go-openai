package openai

import (
	"context"
)

const (
	assistantsSuffix      = "/assistants"
	assistantsFilesSuffix = "/files"
)

type Assistant struct {
	ID             string                 `json:"id"`
	Object         string                 `json:"object"`
	CreatedAt      int64                  `json:"created_at"`
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Model          string                 `json:"model"`
	Instructions   *string                `json:"instructions,omitempty"`
	Tools          []AssistantTool        `json:"tools"`
	ToolResources  *AssistantToolResource `json:"tool_resources,omitempty"`
	FileIDs        []string               `json:"file_ids,omitempty"`
	Metadata       map[string]any         `json:"metadata,omitempty"`
	Temperature    *float32               `json:"temperature,omitempty"`
	TopP           *float32               `json:"top_p,omitempty"`
	ResponseFormat any                    `json:"response_format,omitempty"`

	httpHeader
}

type AssistantToolType string

const (
	AssistantToolTypeCodeInterpreter AssistantToolType = "code_interpreter"
	AssistantToolTypeRetrieval       AssistantToolType = "retrieval"
	AssistantToolTypeFunction        AssistantToolType = "function"
	AssistantToolTypeFileSearch      AssistantToolType = "file_search"
)

type AssistantTool struct {
	Type     AssistantToolType   `json:"type"`
	Function *FunctionDefinition `json:"function,omitempty"`
}

type AssistantToolFileSearch struct {
	VectorStoreIDs []string `json:"vector_store_ids"`
}

type AssistantToolCodeInterpreter struct {
	FileIDs []string `json:"file_ids"`
}

type AssistantToolResource struct {
	FileSearch      *AssistantToolFileSearch      `json:"file_search,omitempty"`
	CodeInterpreter *AssistantToolCodeInterpreter `json:"code_interpreter,omitempty"`
}

type AssistantRequest struct {
	Model          string                 `json:"model"`
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Instructions   *string                `json:"instructions,omitempty"`
	Tools          []AssistantTool        `json:"-"`
	FileIDs        []string               `json:"file_ids,omitempty"`
	Metadata       map[string]any         `json:"metadata,omitempty"`
	ToolResources  *AssistantToolResource `json:"tool_resources,omitempty"`
	ResponseFormat any                    `json:"response_format,omitempty"`
	Temperature    *float32               `json:"temperature,omitempty"`
	TopP           *float32               `json:"top_p,omitempty"`
}

func (a AssistantRequest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type AssistantsList struct {
	Assistants []Assistant `json:"data"`
	LastID     *string     `json:"last_id"`
	FirstID    *string     `json:"first_id"`
	HasMore    bool        `json:"has_more"`
	httpHeader
}

type AssistantDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`

	httpHeader
}

type AssistantFile struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	CreatedAt   int64  `json:"created_at"`
	AssistantID string `json:"assistant_id"`

	httpHeader
}

type AssistantFileRequest struct {
	FileID string `json:"file_id"`
}

type AssistantFilesList struct {
	AssistantFiles []AssistantFile `json:"data"`

	httpHeader
}

func (c *Client) CreateAssistant(ctx context.Context, request AssistantRequest) (response Assistant, err error) {
	_ = "STUB: not implemented"
	return *new(Assistant), nil
}

func (c *Client) RetrieveAssistant(
	ctx context.Context,
	assistantID string,
) (response Assistant, err error) {
	_ = "STUB: not implemented"
	return *new(Assistant), nil
}

func (c *Client) ModifyAssistant(
	ctx context.Context,
	assistantID string,
	request AssistantRequest,
) (response Assistant, err error) {
	_ = "STUB: not implemented"
	return *new(Assistant), nil
}

func (c *Client) DeleteAssistant(
	ctx context.Context,
	assistantID string,
) (response AssistantDeleteResponse, err error) {
	_ = "STUB: not implemented"
	return *new(AssistantDeleteResponse), nil
}

func (c *Client) ListAssistants(
	ctx context.Context,
	limit *int,
	order *string,
	after *string,
	before *string,
) (response AssistantsList, err error) {
	_ = "STUB: not implemented"
	return *new(AssistantsList), nil
}

func (c *Client) CreateAssistantFile(
	ctx context.Context,
	assistantID string,
	request AssistantFileRequest,
) (response AssistantFile, err error) {
	_ = "STUB: not implemented"
	return *new(AssistantFile), nil
}

func (c *Client) RetrieveAssistantFile(
	ctx context.Context,
	assistantID string,
	fileID string,
) (response AssistantFile, err error) {
	_ = "STUB: not implemented"
	return *new(AssistantFile), nil
}

func (c *Client) DeleteAssistantFile(
	ctx context.Context,
	assistantID string,
	fileID string,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ListAssistantFiles(
	ctx context.Context,
	assistantID string,
	limit *int,
	order *string,
	after *string,
	before *string,
) (response AssistantFilesList, err error) {
	_ = "STUB: not implemented"
	return *new(AssistantFilesList), nil
}
