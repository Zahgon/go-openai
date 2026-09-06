package openai

import (
	"context"
	"errors"
)

const (
	ModerationOmniLatest   = "omni-moderation-latest"
	ModerationOmni20240926 = "omni-moderation-2024-09-26"
	ModerationTextStable   = "text-moderation-stable"
	ModerationTextLatest   = "text-moderation-latest"

	ModerationText001 = "text-moderation-001"
)

var (
	ErrModerationInvalidModel = errors.New("this model is not supported with moderation, please use text-moderation-stable or text-moderation-latest instead") //nolint:lll
)

var validModerationModel = map[string]struct{}{
	ModerationOmniLatest:   {},
	ModerationOmni20240926: {},
	ModerationTextStable:   {},
	ModerationTextLatest:   {},
}

type ModerationRequest struct {
	Input string `json:"input,omitempty"`
	Model string `json:"model,omitempty"`
}

type Result struct {
	Categories     ResultCategories     `json:"categories"`
	CategoryScores ResultCategoryScores `json:"category_scores"`
	Flagged        bool                 `json:"flagged"`
}

type ResultCategories struct {
	Hate                  bool `json:"hate"`
	HateThreatening       bool `json:"hate/threatening"`
	Harassment            bool `json:"harassment"`
	HarassmentThreatening bool `json:"harassment/threatening"`
	SelfHarm              bool `json:"self-harm"`
	SelfHarmIntent        bool `json:"self-harm/intent"`
	SelfHarmInstructions  bool `json:"self-harm/instructions"`
	Sexual                bool `json:"sexual"`
	SexualMinors          bool `json:"sexual/minors"`
	Violence              bool `json:"violence"`
	ViolenceGraphic       bool `json:"violence/graphic"`
}

type ResultCategoryScores struct {
	Hate                  float32 `json:"hate"`
	HateThreatening       float32 `json:"hate/threatening"`
	Harassment            float32 `json:"harassment"`
	HarassmentThreatening float32 `json:"harassment/threatening"`
	SelfHarm              float32 `json:"self-harm"`
	SelfHarmIntent        float32 `json:"self-harm/intent"`
	SelfHarmInstructions  float32 `json:"self-harm/instructions"`
	Sexual                float32 `json:"sexual"`
	SexualMinors          float32 `json:"sexual/minors"`
	Violence              float32 `json:"violence"`
	ViolenceGraphic       float32 `json:"violence/graphic"`
}

type ModerationResponse struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Results []Result `json:"results"`

	httpHeader
}

func (c *Client) Moderations(ctx context.Context, request ModerationRequest) (response ModerationResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ModerationResponse), nil
}
