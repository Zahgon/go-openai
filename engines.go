package openai

import (
	"context"
)

type Engine struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`

	httpHeader
}

type EnginesList struct {
	Engines []Engine `json:"data"`

	httpHeader
}

func (c *Client) ListEngines(ctx context.Context) (engines EnginesList, err error) {
	_ = "STUB: not implemented"
	return *new(EnginesList), nil
}

func (c *Client) GetEngine(
	ctx context.Context,
	engineID string,
) (engine Engine, err error) {
	_ = "STUB: not implemented"
	return *new(Engine), nil
}
