package openai

import (
	"context"
)

type FineTuningJob struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	CreatedAt       int64           `json:"created_at"`
	FinishedAt      int64           `json:"finished_at"`
	Model           string          `json:"model"`
	FineTunedModel  string          `json:"fine_tuned_model,omitempty"`
	OrganizationID  string          `json:"organization_id"`
	Status          string          `json:"status"`
	Hyperparameters Hyperparameters `json:"hyperparameters"`
	TrainingFile    string          `json:"training_file"`
	ValidationFile  string          `json:"validation_file,omitempty"`
	ResultFiles     []string        `json:"result_files"`
	TrainedTokens   int             `json:"trained_tokens"`

	httpHeader
}

type Hyperparameters struct {
	Epochs                 any `json:"n_epochs,omitempty"`
	LearningRateMultiplier any `json:"learning_rate_multiplier,omitempty"`
	BatchSize              any `json:"batch_size,omitempty"`
}

type FineTuningJobRequest struct {
	TrainingFile    string           `json:"training_file"`
	ValidationFile  string           `json:"validation_file,omitempty"`
	Model           string           `json:"model,omitempty"`
	Hyperparameters *Hyperparameters `json:"hyperparameters,omitempty"`
	Suffix          string           `json:"suffix,omitempty"`
}

type FineTuningJobEventList struct {
	Object  string          `json:"object"`
	Data    []FineTuneEvent `json:"data"`
	HasMore bool            `json:"has_more"`

	httpHeader
}

type FineTuningJobEvent struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	CreatedAt int    `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
	Type      string `json:"type"`
}

func (c *Client) CreateFineTuningJob(
	ctx context.Context,
	request FineTuningJobRequest,
) (response FineTuningJob, err error) {
	_ = "STUB: not implemented"
	return *new(FineTuningJob), nil
}

func (c *Client) CancelFineTuningJob(ctx context.Context, fineTuningJobID string) (response FineTuningJob, err error) {
	_ = "STUB: not implemented"
	return *new(FineTuningJob), nil
}

func (c *Client) RetrieveFineTuningJob(
	ctx context.Context,
	fineTuningJobID string,
) (response FineTuningJob, err error) {
	_ = "STUB: not implemented"
	return *new(FineTuningJob), nil
}

type listFineTuningJobEventsParameters struct {
	after *string
	limit *int
}

type ListFineTuningJobEventsParameter func(*listFineTuningJobEventsParameters)

func ListFineTuningJobEventsWithAfter(after string) ListFineTuningJobEventsParameter {
	_ = "STUB: not implemented"
	return *new(ListFineTuningJobEventsParameter)
}

func ListFineTuningJobEventsWithLimit(limit int) ListFineTuningJobEventsParameter {
	_ = "STUB: not implemented"
	return *new(ListFineTuningJobEventsParameter)
}

func (c *Client) ListFineTuningJobEvents(
	ctx context.Context,
	fineTuningJobID string,
	setters ...ListFineTuningJobEventsParameter,
) (response FineTuningJobEventList, err error) {
	_ = "STUB: not implemented"
	return *new(FineTuningJobEventList), nil
}
