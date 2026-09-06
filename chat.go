package openai

import (
	"context"
	"encoding/json"
	"errors"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleFunction  = "function"
	ChatMessageRoleTool      = "tool"
	ChatMessageRoleDeveloper = "developer"
)

const chatCompletionsSuffix = "/chat/completions"

var (
	ErrChatCompletionInvalidModel       = errors.New("this model is not supported with this method, please use CreateCompletion client method instead") //nolint:lll
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream")              //nolint:lll
	ErrContentFieldsMisused             = errors.New("can't use both Content and MultiContent properties simultaneously")
)

type Hate struct {
	Filtered bool   `json:"filtered"`
	Severity string `json:"severity,omitempty"`
}
type SelfHarm struct {
	Filtered bool   `json:"filtered"`
	Severity string `json:"severity,omitempty"`
}
type Sexual struct {
	Filtered bool   `json:"filtered"`
	Severity string `json:"severity,omitempty"`
}
type Violence struct {
	Filtered bool   `json:"filtered"`
	Severity string `json:"severity,omitempty"`
}

type JailBreak struct {
	Filtered bool `json:"filtered"`
	Detected bool `json:"detected"`
}

type Profanity struct {
	Filtered bool `json:"filtered"`
	Detected bool `json:"detected"`
}

type ContentFilterResults struct {
	Hate      Hate      `json:"hate,omitempty"`
	SelfHarm  SelfHarm  `json:"self_harm,omitempty"`
	Sexual    Sexual    `json:"sexual,omitempty"`
	Violence  Violence  `json:"violence,omitempty"`
	JailBreak JailBreak `json:"jailbreak,omitempty"`
	Profanity Profanity `json:"profanity,omitempty"`
}

type PromptAnnotation struct {
	PromptIndex          int                  `json:"prompt_index,omitempty"`
	ContentFilterResults ContentFilterResults `json:"content_filter_results,omitempty"`
}

type ImageURLDetail string

const (
	ImageURLDetailHigh ImageURLDetail = "high"
	ImageURLDetailLow  ImageURLDetail = "low"
	ImageURLDetailAuto ImageURLDetail = "auto"
)

type ChatMessageImageURL struct {
	URL    string         `json:"url,omitempty"`
	Detail ImageURLDetail `json:"detail,omitempty"`
}

type ChatMessagePartType string

const (
	ChatMessagePartTypeText     ChatMessagePartType = "text"
	ChatMessagePartTypeImageURL ChatMessagePartType = "image_url"
)

type ChatMessagePart struct {
	Type     ChatMessagePartType  `json:"type,omitempty"`
	Text     string               `json:"text,omitempty"`
	ImageURL *ChatMessageImageURL `json:"image_url,omitempty"`
}

type ChatCompletionMessage struct {
	Role         string `json:"role"`
	Content      string `json:"content,omitempty"`
	Refusal      string `json:"refusal,omitempty"`
	MultiContent []ChatMessagePart

	Name string `json:"name,omitempty"`

	ReasoningContent string `json:"reasoning_content,omitempty"`

	FunctionCall *FunctionCall `json:"function_call,omitempty"`

	ToolCalls []ToolCall `json:"tool_calls,omitempty"`

	ToolCallID string `json:"tool_call_id,omitempty"`
}

func (m ChatCompletionMessage) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ChatCompletionMessage) UnmarshalJSON(bs []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type ToolCall struct {
	Index    *int         `json:"index,omitempty"`
	ID       string       `json:"id,omitempty"`
	Type     ToolType     `json:"type"`
	Function FunctionCall `json:"function"`
}

type FunctionCall struct {
	Name string `json:"name,omitempty"`

	Arguments string `json:"arguments,omitempty"`
}

type ChatCompletionResponseFormatType string

const (
	ChatCompletionResponseFormatTypeJSONObject ChatCompletionResponseFormatType = "json_object"
	ChatCompletionResponseFormatTypeJSONSchema ChatCompletionResponseFormatType = "json_schema"
	ChatCompletionResponseFormatTypeText       ChatCompletionResponseFormatType = "text"
)

type ChatCompletionResponseFormat struct {
	Type       ChatCompletionResponseFormatType        `json:"type,omitempty"`
	JSONSchema *ChatCompletionResponseFormatJSONSchema `json:"json_schema,omitempty"`
}

type ChatCompletionResponseFormatJSONSchema struct {
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Schema      json.Marshaler `json:"schema"`
	Strict      bool           `json:"strict"`
}

func (r *ChatCompletionResponseFormatJSONSchema) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type ChatCompletionRequestExtensions struct {
	GuidedChoice []string `json:"guided_choice,omitempty"`
}

type ChatCompletionRequest struct {
	Model    string                  `json:"model"`
	Messages []ChatCompletionMessage `json:"messages"`

	MaxTokens int `json:"max_tokens,omitempty"`

	MaxCompletionTokens int                           `json:"max_completion_tokens,omitempty"`
	Temperature         float32                       `json:"temperature,omitempty"`
	TopP                float32                       `json:"top_p,omitempty"`
	N                   int                           `json:"n,omitempty"`
	Stream              bool                          `json:"stream"`
	Stop                []string                      `json:"stop,omitempty"`
	PresencePenalty     float32                       `json:"presence_penalty,omitempty"`
	ResponseFormat      *ChatCompletionResponseFormat `json:"response_format,omitempty"`
	Seed                *int                          `json:"seed,omitempty"`
	FrequencyPenalty    float32                       `json:"frequency_penalty,omitempty"`

	LogitBias map[string]int `json:"logit_bias,omitempty"`

	LogProbs bool `json:"logprobs,omitempty"`

	TopLogProbs int    `json:"top_logprobs,omitempty"`
	User        string `json:"user,omitempty"`

	Functions []FunctionDefinition `json:"functions,omitempty"`

	FunctionCall any    `json:"function_call,omitempty"`
	Tools        []Tool `json:"tools,omitempty"`

	ToolChoice any `json:"tool_choice,omitempty"`

	StreamOptions *StreamOptions `json:"stream_options,omitempty"`

	ParallelToolCalls any `json:"parallel_tool_calls,omitempty"`

	Store bool `json:"store,omitempty"`

	ReasoningEffort string `json:"reasoning_effort,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	Prediction *Prediction `json:"prediction,omitempty"`

	ChatTemplateKwargs map[string]any `json:"chat_template_kwargs,omitempty"`

	ServiceTier ServiceTier `json:"service_tier,omitempty"`

	Verbosity string `json:"verbosity,omitempty"`

	SafetyIdentifier string `json:"safety_identifier,omitempty"`

	ChatCompletionRequestExtensions
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage,omitempty"`
}

type ToolType string

const (
	ToolTypeFunction           ToolType = "function"
	ToolTypeWebSearch          ToolType = "web_search"
	ToolTypeWebSearchPreview   ToolType = "web_search_preview"
	ToolTypeFileSearch         ToolType = "file_search"
	ToolTypeComputer           ToolType = "computer"
	ToolTypeComputerUsePreview ToolType = "computer_use_preview"
	ToolTypeComputerUse        ToolType = "computer_use"
	ToolTypeCodeInterpreter    ToolType = "code_interpreter"
	ToolTypeImageGeneration    ToolType = "image_generation"
	ToolTypeMCP                ToolType = "mcp"
	ToolTypeCustom             ToolType = "custom"
	ToolTypeLocalShell         ToolType = "local_shell"
	ToolTypeShell              ToolType = "shell"
	ToolTypeApplyPatch         ToolType = "apply_patch"
	ToolTypeToolSearch         ToolType = "tool_search"
)

type Tool struct {
	Type     ToolType            `json:"type"`
	Function *FunctionDefinition `json:"function,omitempty"`

	Parameters map[string]any `json:"-"`
}

func (t Tool) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type ToolChoice struct {
	Type     ToolType     `json:"type"`
	Function ToolFunction `json:"function,omitempty"`
}

type ToolFunction struct {
	Name string `json:"name"`
}

type FunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Strict      bool   `json:"strict,omitempty"`

	Parameters any `json:"parameters"`
}

type FunctionDefine = FunctionDefinition

type TopLogProbs struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes,omitempty"`
}

type LogProb struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes,omitempty"`

	TopLogProbs []TopLogProbs `json:"top_logprobs"`
}

type LogProbs struct {
	Content []LogProb `json:"content"`
}

type Prediction struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type FinishReason string

const (
	FinishReasonStop          FinishReason = "stop"
	FinishReasonLength        FinishReason = "length"
	FinishReasonFunctionCall  FinishReason = "function_call"
	FinishReasonToolCalls     FinishReason = "tool_calls"
	FinishReasonContentFilter FinishReason = "content_filter"
	FinishReasonNull          FinishReason = "null"
)

type ServiceTier string

const (
	ServiceTierAuto     ServiceTier = "auto"
	ServiceTierDefault  ServiceTier = "default"
	ServiceTierFlex     ServiceTier = "flex"
	ServiceTierPriority ServiceTier = "priority"
)

func (r FinishReason) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type ChatCompletionChoice struct {
	Index   int                   `json:"index"`
	Message ChatCompletionMessage `json:"message"`

	FinishReason         FinishReason         `json:"finish_reason"`
	LogProbs             *LogProbs            `json:"logprobs,omitempty"`
	ContentFilterResults ContentFilterResults `json:"content_filter_results,omitempty"`
}

type ChatCompletionResponse struct {
	ID                  string                 `json:"id"`
	Object              string                 `json:"object"`
	Created             int64                  `json:"created"`
	Model               string                 `json:"model"`
	Choices             []ChatCompletionChoice `json:"choices"`
	Usage               Usage                  `json:"usage"`
	SystemFingerprint   string                 `json:"system_fingerprint"`
	PromptFilterResults []PromptFilterResult   `json:"prompt_filter_results,omitempty"`
	ServiceTier         ServiceTier            `json:"service_tier,omitempty"`

	httpHeader
}

func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request ChatCompletionRequest,
) (response ChatCompletionResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ChatCompletionResponse), nil
}
