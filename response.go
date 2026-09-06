package openai

import (
	"context"
	"errors"
	"net/url"
)

const responsesSuffix = "/responses"

var ErrResponseStreamNotSupported = errors.New(
	"streaming is not supported with this method, please use CreateResponseStream",
)

type ResponseInclude string

const (
	ResponseIncludeFileSearchCallResults      ResponseInclude = "file_search_call.results"
	ResponseIncludeWebSearchCallResults       ResponseInclude = "web_search_call.results"
	ResponseIncludeWebSearchCallActionSources ResponseInclude = "web_search_call.action.sources"
	ResponseIncludeInputImageURL              ResponseInclude = "message.input_image.image_url"
	ResponseIncludeComputerCallOutputImageURL ResponseInclude = "computer_call_output.output.image_url"
	ResponseIncludeCodeInterpreterCallOutputs ResponseInclude = "code_interpreter_call.outputs"
	ResponseIncludeReasoningEncryptedContent  ResponseInclude = "reasoning.encrypted_content"
	ResponseIncludeMessageOutputTextLogprobs  ResponseInclude = "message.output_text.logprobs"
)

type ResponseStatus string

const (
	ResponseStatusQueued     ResponseStatus = "queued"
	ResponseStatusInProgress ResponseStatus = "in_progress"
	ResponseStatusCompleted  ResponseStatus = "completed"
	ResponseStatusFailed     ResponseStatus = "failed"
	ResponseStatusIncomplete ResponseStatus = "incomplete"
	ResponseStatusCancelling ResponseStatus = "cancelling"
	ResponseStatusCancelled  ResponseStatus = "cancelled"
)

type ResponseTruncation string

const (
	ResponseTruncationAuto     ResponseTruncation = "auto"
	ResponseTruncationDisabled ResponseTruncation = "disabled"
)

type ResponseTool = Tool

func NewResponseFunctionTool(function FunctionDefinition) ResponseTool {
	_ = "STUB: not implemented"
	return *new(ResponseTool)
}

type ResponseReasoning struct {
	Effort          string `json:"effort,omitempty"`
	GenerateSummary string `json:"generate_summary,omitempty"`
	Summary         string `json:"summary,omitempty"`
	Context         string `json:"context,omitempty"`
	Mode            string `json:"mode,omitempty"`
}

type ResponseStreamOptions struct {
	IncludeObfuscation *bool `json:"include_obfuscation,omitempty"`
}

type ResponseTextConfig struct {
	Format    *ResponseTextFormat `json:"format,omitempty"`
	Verbosity string              `json:"verbosity,omitempty"`
}

type ResponseTextFormat struct {
	Type        string `json:"type"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Schema      any    `json:"schema,omitempty"`
	Strict      bool   `json:"strict,omitempty"`
}

type ResponsePrompt struct {
	ID        string         `json:"id"`
	Variables map[string]any `json:"variables,omitempty"`
	Version   string         `json:"version,omitempty"`
}

type ResponsePromptCacheOptions struct {
	Mode string `json:"mode,omitempty"`
	TTL  string `json:"ttl,omitempty"`
}

type ResponsePromptCacheBreakpoint struct {
	Mode string `json:"mode"`
}

type CreateResponseRequest struct {
	Background           bool                        `json:"background,omitempty"`
	ContextManagement    []any                       `json:"context_management,omitempty"`
	Conversation         any                         `json:"conversation,omitempty"`
	Include              []ResponseInclude           `json:"include,omitempty"`
	Input                any                         `json:"input"`
	Instructions         string                      `json:"instructions,omitempty"`
	MaxOutputTokens      int                         `json:"max_output_tokens,omitempty"`
	MaxToolCalls         int                         `json:"max_tool_calls,omitempty"`
	Metadata             map[string]any              `json:"metadata,omitempty"`
	Model                string                      `json:"model,omitempty"`
	Moderation           any                         `json:"moderation,omitempty"`
	ParallelToolCalls    *bool                       `json:"parallel_tool_calls,omitempty"`
	PreviousResponseID   string                      `json:"previous_response_id,omitempty"`
	Prompt               *ResponsePrompt             `json:"prompt,omitempty"`
	PromptCacheKey       string                      `json:"prompt_cache_key,omitempty"`
	PromptCacheOptions   *ResponsePromptCacheOptions `json:"prompt_cache_options,omitempty"`
	PromptCacheRetention string                      `json:"prompt_cache_retention,omitempty"`
	Reasoning            *ResponseReasoning          `json:"reasoning,omitempty"`
	SafetyIdentifier     string                      `json:"safety_identifier,omitempty"`
	ServiceTier          string                      `json:"service_tier,omitempty"`
	Store                *bool                       `json:"store,omitempty"`
	Stream               bool                        `json:"stream,omitempty"`
	StreamOptions        *ResponseStreamOptions      `json:"stream_options,omitempty"`
	Temperature          *float32                    `json:"temperature,omitempty"`
	Text                 *ResponseTextConfig         `json:"text,omitempty"`
	ToolChoice           any                         `json:"tool_choice,omitempty"`
	Tools                []ResponseTool              `json:"tools,omitempty"`
	TopLogprobs          int                         `json:"top_logprobs,omitempty"`
	TopP                 *float32                    `json:"top_p,omitempty"`
	Truncation           ResponseTruncation          `json:"truncation,omitempty"`
	User                 string                      `json:"user,omitempty"`
	ExtraBody            map[string]any              `json:"-"`
}

func (r CreateResponseRequest) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResponseInputMessage struct {
	Type    string `json:"type,omitempty"`
	Role    string `json:"role"`
	Content any    `json:"content"`
	Status  string `json:"status,omitempty"`
	Phase   string `json:"phase,omitempty"`
}

type ResponseInputText struct {
	Type                  string                         `json:"type"`
	Text                  string                         `json:"text"`
	PromptCacheBreakpoint *ResponsePromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

type ResponseInputImage struct {
	Type                  string                         `json:"type"`
	Detail                string                         `json:"detail,omitempty"`
	FileID                string                         `json:"file_id,omitempty"`
	ImageURL              string                         `json:"image_url,omitempty"`
	PromptCacheBreakpoint *ResponsePromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

type ResponseInputFile struct {
	Type                  string                         `json:"type"`
	FileData              string                         `json:"file_data,omitempty"`
	FileID                string                         `json:"file_id,omitempty"`
	FileURL               string                         `json:"file_url,omitempty"`
	Filename              string                         `json:"filename,omitempty"`
	Detail                string                         `json:"detail,omitempty"`
	PromptCacheBreakpoint *ResponsePromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

type ResponseFunctionCallOutput struct {
	Type   string `json:"type"`
	CallID string `json:"call_id"`
	Output any    `json:"output"`
	Status string `json:"status,omitempty"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseIncompleteDetails struct {
	Reason string `json:"reason,omitempty"`
}

type ResponseConversation struct {
	ID string `json:"id"`
}

type ResponseUsage struct {
	InputTokens         int                          `json:"input_tokens"`
	InputTokensDetails  *ResponseInputTokensDetails  `json:"input_tokens_details,omitempty"`
	OutputTokens        int                          `json:"output_tokens"`
	OutputTokensDetails *ResponseOutputTokensDetails `json:"output_tokens_details,omitempty"`
	TotalTokens         int                          `json:"total_tokens"`
}

type ResponseInputTokensDetails struct {
	CachedTokens     int `json:"cached_tokens"`
	CacheWriteTokens int `json:"cache_write_tokens"`
}

type ResponseOutputTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type ResponseAnnotation struct {
	Type       string `json:"type"`
	FileID     string `json:"file_id,omitempty"`
	Filename   string `json:"filename,omitempty"`
	Index      int    `json:"index,omitempty"`
	StartIndex int    `json:"start_index,omitempty"`
	EndIndex   int    `json:"end_index,omitempty"`
	URL        string `json:"url,omitempty"`
	Title      string `json:"title,omitempty"`
}

type ResponseLogprob struct {
	Token       string            `json:"token"`
	Bytes       []int64           `json:"bytes,omitempty"`
	Logprob     float64           `json:"logprob"`
	TopLogprobs []ResponseLogprob `json:"top_logprobs,omitempty"`
}

type ResponseOutputContent struct {
	Type        string               `json:"type"`
	Text        string               `json:"text,omitempty"`
	Refusal     string               `json:"refusal,omitempty"`
	Annotations []ResponseAnnotation `json:"annotations,omitempty"`
	Logprobs    []ResponseLogprob    `json:"logprobs,omitempty"`
}

type ResponseSummaryPart struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ResponseOutputItem struct {
	ID        string                  `json:"id,omitempty"`
	Type      string                  `json:"type"`
	Status    string                  `json:"status,omitempty"`
	Role      string                  `json:"role,omitempty"`
	Content   []ResponseOutputContent `json:"content,omitempty"`
	CallID    string                  `json:"call_id,omitempty"`
	Name      string                  `json:"name,omitempty"`
	Arguments string                  `json:"arguments,omitempty"`
	Summary   []ResponseSummaryPart   `json:"summary,omitempty"`
	Action    any                     `json:"action,omitempty"`
	Results   any                     `json:"results,omitempty"`
	Output    any                     `json:"output,omitempty"`
}

type CreateResponseResponse struct {
	ID                   string                      `json:"id"`
	Object               string                      `json:"object"`
	Created              int64                       `json:"created_at"`
	CompletedAt          *int64                      `json:"completed_at,omitempty"`
	Status               ResponseStatus              `json:"status,omitempty"`
	Error                *ResponseError              `json:"error,omitempty"`
	IncompleteDetails    *ResponseIncompleteDetails  `json:"incomplete_details,omitempty"`
	Instructions         any                         `json:"instructions,omitempty"`
	MaxOutputTokens      *int                        `json:"max_output_tokens,omitempty"`
	MaxToolCalls         *int                        `json:"max_tool_calls,omitempty"`
	Metadata             map[string]any              `json:"metadata,omitempty"`
	Model                string                      `json:"model"`
	Moderation           any                         `json:"moderation,omitempty"`
	Output               []any                       `json:"output"`
	OutputText           string                      `json:"output_text,omitempty"`
	ParallelToolCalls    bool                        `json:"parallel_tool_calls,omitempty"`
	PreviousResponseID   string                      `json:"previous_response_id,omitempty"`
	Reasoning            *ResponseReasoning          `json:"reasoning,omitempty"`
	ServiceTier          string                      `json:"service_tier,omitempty"`
	Store                bool                        `json:"store,omitempty"`
	Temperature          *float32                    `json:"temperature,omitempty"`
	Text                 *ResponseTextConfig         `json:"text,omitempty"`
	ToolChoice           any                         `json:"tool_choice,omitempty"`
	Tools                []any                       `json:"tools,omitempty"`
	TopLogprobs          int                         `json:"top_logprobs,omitempty"`
	TopP                 *float32                    `json:"top_p,omitempty"`
	Truncation           ResponseTruncation          `json:"truncation,omitempty"`
	Usage                *ResponseUsage              `json:"usage,omitempty"`
	Background           *bool                       `json:"background,omitempty"`
	Conversation         *ResponseConversation       `json:"conversation,omitempty"`
	Prompt               *ResponsePrompt             `json:"prompt,omitempty"`
	PromptCacheKey       string                      `json:"prompt_cache_key,omitempty"`
	PromptCacheOptions   *ResponsePromptCacheOptions `json:"prompt_cache_options,omitempty"`
	PromptCacheRetention string                      `json:"prompt_cache_retention,omitempty"`
	SafetyIdentifier     string                      `json:"safety_identifier,omitempty"`
	User                 string                      `json:"user,omitempty"`

	httpHeader
}

func (r CreateResponseResponse) GetOutputText() string { _ = "STUB: not implemented"; return "" }

type RetrieveResponseOptions struct {
	Include            []ResponseInclude
	IncludeObfuscation *bool
	StartingAfter      *int
}

type ResponseInputItemsListOptions struct {
	After   string
	Include []ResponseInclude
	Limit   int
	Order   string
}

type ResponseInputItemsList struct {
	Object  string `json:"object"`
	Data    []any  `json:"data"`
	FirstID string `json:"first_id"`
	LastID  string `json:"last_id"`
	HasMore bool   `json:"has_more"`

	httpHeader
}

type ResponseDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`

	httpHeader
}

type DeleteResponseResponse = ResponseDeleteResponse

type ResponseInputTokensResponse struct {
	Object      string `json:"object"`
	InputTokens int    `json:"input_tokens"`

	httpHeader
}

type ResponseInputTokensRequest = CreateResponseRequest

type ResponseCompaction struct {
	ID        string         `json:"id"`
	Object    string         `json:"object"`
	CreatedAt int64          `json:"created_at"`
	Output    []any          `json:"output"`
	Usage     *ResponseUsage `json:"usage,omitempty"`

	httpHeader
}

type CompactResponseRequest = CreateResponseRequest

func (c *Client) CreateResponse(
	ctx context.Context,
	request CreateResponseRequest,
) (response CreateResponseResponse, err error) {
	_ = "STUB: not implemented"
	return *new(CreateResponseResponse), nil
}

func (c *Client) RetrieveResponse(
	ctx context.Context,
	responseID string,
	options ...RetrieveResponseOptions,
) (response CreateResponseResponse, err error) {
	_ = "STUB: not implemented"
	return *new(CreateResponseResponse), nil
}

func (c *Client) GetResponse(
	ctx context.Context,
	responseID string,
	options ...RetrieveResponseOptions,
) (CreateResponseResponse, error) {
	_ = "STUB: not implemented"
	return *new(CreateResponseResponse), nil
}

func (c *Client) DeleteResponse(ctx context.Context, responseID string) (response ResponseDeleteResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ResponseDeleteResponse), nil
}

func (c *Client) CancelResponse(ctx context.Context, responseID string) (response CreateResponseResponse, err error) {
	_ = "STUB: not implemented"
	return *new(CreateResponseResponse), nil
}

func (c *Client) ListResponseInputItems(
	ctx context.Context,
	responseID string,
	options ...ResponseInputItemsListOptions,
) (response ResponseInputItemsList, err error) {
	_ = "STUB: not implemented"
	return *new(ResponseInputItemsList), nil
}

func (c *Client) CountResponseInputTokens(
	ctx context.Context,
	request ResponseInputTokensRequest,
) (response ResponseInputTokensResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ResponseInputTokensResponse), nil
}

func (c *Client) CompactResponse(
	ctx context.Context,
	request CompactResponseRequest,
) (response ResponseCompaction, err error) {
	_ = "STUB: not implemented"
	return *new(ResponseCompaction), nil
}

func responseResourceSuffix(responseID, action string, values url.Values) string {
	_ = "STUB: not implemented"
	return ""
}
