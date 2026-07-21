package openai

import (
	"context"
)

type Run struct {
	ID             string             `json:"id"`
	Object         string             `json:"object"`
	CreatedAt      int64              `json:"created_at"`
	ThreadID       string             `json:"thread_id"`
	AssistantID    string             `json:"assistant_id"`
	Status         RunStatus          `json:"status"`
	RequiredAction *RunRequiredAction `json:"required_action,omitempty"`
	LastError      *RunLastError      `json:"last_error,omitempty"`
	ExpiresAt      int64              `json:"expires_at"`
	StartedAt      *int64             `json:"started_at,omitempty"`
	CancelledAt    *int64             `json:"cancelled_at,omitempty"`
	FailedAt       *int64             `json:"failed_at,omitempty"`
	CompletedAt    *int64             `json:"completed_at,omitempty"`
	Model          string             `json:"model"`
	Instructions   string             `json:"instructions,omitempty"`
	Tools          []Tool             `json:"tools"`
	FileIDS        []string           `json:"file_ids"` //nolint:revive // backwards-compatibility
	Metadata       map[string]any     `json:"metadata"`
	Usage          Usage              `json:"usage,omitempty"`

	Temperature *float32 `json:"temperature,omitempty"`

	MaxPromptTokens int `json:"max_prompt_tokens,omitempty"`

	MaxCompletionTokens int `json:"max_completion_tokens,omitempty"`

	TruncationStrategy *ThreadTruncationStrategy `json:"truncation_strategy,omitempty"`

	httpHeader
}

type RunStatus string

const (
	RunStatusQueued         RunStatus = "queued"
	RunStatusInProgress     RunStatus = "in_progress"
	RunStatusRequiresAction RunStatus = "requires_action"
	RunStatusCancelling     RunStatus = "cancelling"
	RunStatusFailed         RunStatus = "failed"
	RunStatusCompleted      RunStatus = "completed"
	RunStatusIncomplete     RunStatus = "incomplete"
	RunStatusExpired        RunStatus = "expired"
	RunStatusCancelled      RunStatus = "cancelled"
)

type RunRequiredAction struct {
	Type              RequiredActionType `json:"type"`
	SubmitToolOutputs *SubmitToolOutputs `json:"submit_tool_outputs,omitempty"`
}

type RequiredActionType string

const (
	RequiredActionTypeSubmitToolOutputs RequiredActionType = "submit_tool_outputs"
)

type SubmitToolOutputs struct {
	ToolCalls []ToolCall `json:"tool_calls"`
}

type RunLastError struct {
	Code    RunError `json:"code"`
	Message string   `json:"message"`
}

type RunError string

const (
	RunErrorServerError       RunError = "server_error"
	RunErrorRateLimitExceeded RunError = "rate_limit_exceeded"
)

type RunRequest struct {
	AssistantID            string          `json:"assistant_id"`
	Model                  string          `json:"model,omitempty"`
	Instructions           string          `json:"instructions,omitempty"`
	AdditionalInstructions string          `json:"additional_instructions,omitempty"`
	AdditionalMessages     []ThreadMessage `json:"additional_messages,omitempty"`
	Tools                  []Tool          `json:"tools,omitempty"`
	Metadata               map[string]any  `json:"metadata,omitempty"`

	Temperature *float32 `json:"temperature,omitempty"`
	TopP        *float32 `json:"top_p,omitempty"`

	MaxPromptTokens int `json:"max_prompt_tokens,omitempty"`

	MaxCompletionTokens int `json:"max_completion_tokens,omitempty"`

	TruncationStrategy *ThreadTruncationStrategy `json:"truncation_strategy,omitempty"`

	ToolChoice any `json:"tool_choice,omitempty"`

	ResponseFormat any `json:"response_format,omitempty"`

	ParallelToolCalls any `json:"parallel_tool_calls,omitempty"`
}

type ThreadTruncationStrategy struct {
	Type TruncationStrategy `json:"type,omitempty"`

	LastMessages *int `json:"last_messages,omitempty"`
}

type TruncationStrategy string

const (
	TruncationStrategyAuto = TruncationStrategy("auto")

	TruncationStrategyLastMessages = TruncationStrategy("last_messages")
)

type ReponseFormat struct {
	Type string `json:"type"`
}

type RunModifyRequest struct {
	Metadata map[string]any `json:"metadata,omitempty"`
}

type RunList struct {
	Runs []Run `json:"data"`

	httpHeader
}

type SubmitToolOutputsRequest struct {
	ToolOutputs []ToolOutput `json:"tool_outputs"`
}

type ToolOutput struct {
	ToolCallID string `json:"tool_call_id"`
	Output     any    `json:"output"`
}

type CreateThreadAndRunRequest struct {
	RunRequest
	Thread ThreadRequest `json:"thread"`
}

type RunStep struct {
	ID          string         `json:"id"`
	Object      string         `json:"object"`
	CreatedAt   int64          `json:"created_at"`
	AssistantID string         `json:"assistant_id"`
	ThreadID    string         `json:"thread_id"`
	RunID       string         `json:"run_id"`
	Type        RunStepType    `json:"type"`
	Status      RunStepStatus  `json:"status"`
	StepDetails StepDetails    `json:"step_details"`
	LastError   *RunLastError  `json:"last_error,omitempty"`
	ExpiredAt   *int64         `json:"expired_at,omitempty"`
	CancelledAt *int64         `json:"cancelled_at,omitempty"`
	FailedAt    *int64         `json:"failed_at,omitempty"`
	CompletedAt *int64         `json:"completed_at,omitempty"`
	Metadata    map[string]any `json:"metadata"`

	httpHeader
}

type RunStepStatus string

const (
	RunStepStatusInProgress RunStepStatus = "in_progress"
	RunStepStatusCancelling RunStepStatus = "cancelled"
	RunStepStatusFailed     RunStepStatus = "failed"
	RunStepStatusCompleted  RunStepStatus = "completed"
	RunStepStatusExpired    RunStepStatus = "expired"
)

type RunStepType string

const (
	RunStepTypeMessageCreation RunStepType = "message_creation"
	RunStepTypeToolCalls       RunStepType = "tool_calls"
)

type StepDetails struct {
	Type            RunStepType                 `json:"type"`
	MessageCreation *StepDetailsMessageCreation `json:"message_creation,omitempty"`
	ToolCalls       []ToolCall                  `json:"tool_calls,omitempty"`
}

type StepDetailsMessageCreation struct {
	MessageID string `json:"message_id"`
}

type RunStepList struct {
	RunSteps []RunStep `json:"data"`

	FirstID string `json:"first_id"`
	LastID  string `json:"last_id"`
	HasMore bool   `json:"has_more"`

	httpHeader
}

type Pagination struct {
	Limit  *int
	Order  *string
	After  *string
	Before *string
}

func (c *Client) CreateRun(
	ctx context.Context,
	threadID string,
	request RunRequest,
) (response Run, err error) {
	_ = "STUB: not implemented"
	return *new(Run), nil
}

func (c *Client) RetrieveRun(
	ctx context.Context,
	threadID string,
	runID string,
) (response Run, err error) {
	_ = "STUB: not implemented"
	return *new(Run), nil
}

func (c *Client) ModifyRun(
	ctx context.Context,
	threadID string,
	runID string,
	request RunModifyRequest,
) (response Run, err error) {
	_ = "STUB: not implemented"
	return *new(Run), nil
}

func (c *Client) ListRuns(
	ctx context.Context,
	threadID string,
	pagination Pagination,
) (response RunList, err error) {
	_ = "STUB: not implemented"
	return *new(RunList), nil
}

func (c *Client) SubmitToolOutputs(
	ctx context.Context,
	threadID string,
	runID string,
	request SubmitToolOutputsRequest) (response Run, err error) {
	_ = "STUB: not implemented"
	return *new(Run), nil
}

func (c *Client) CancelRun(
	ctx context.Context,
	threadID string,
	runID string) (response Run, err error) {
	_ = "STUB: not implemented"
	return *new(Run), nil
}

func (c *Client) CreateThreadAndRun(
	ctx context.Context,
	request CreateThreadAndRunRequest) (response Run, err error) {
	_ = "STUB: not implemented"
	return *new(Run), nil
}

func (c *Client) RetrieveRunStep(
	ctx context.Context,
	threadID string,
	runID string,
	stepID string,
) (response RunStep, err error) {
	_ = "STUB: not implemented"
	return *new(RunStep), nil
}

func (c *Client) ListRunSteps(
	ctx context.Context,
	threadID string,
	runID string,
	pagination Pagination,
) (response RunStepList, err error) {
	_ = "STUB: not implemented"
	return *new(RunStepList), nil
}
