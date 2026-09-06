package openai

import (
	"context"
)

type Model struct {
	CreatedAt  int64        `json:"created"`
	ID         string       `json:"id"`
	Object     string       `json:"object"`
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
	Root       string       `json:"root"`
	Parent     string       `json:"parent"`

	httpHeader
}

type Permission struct {
	CreatedAt          int64       `json:"created"`
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}

type FineTuneModelDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`

	httpHeader
}

type ModelsList struct {
	Models []Model `json:"data"`

	httpHeader
}

func (c *Client) ListModels(ctx context.Context) (models ModelsList, err error) {
	_ = "STUB: not implemented"
	return *new(ModelsList), nil
}

func (c *Client) GetModel(ctx context.Context, modelID string) (model Model, err error) {
	_ = "STUB: not implemented"
	return *new(Model), nil
}

func (c *Client) DeleteFineTuneModel(ctx context.Context, modelID string) (
	response FineTuneModelDeleteResponse, err error) {
	_ = "STUB: not implemented"
	return *new(FineTuneModelDeleteResponse), nil
}
