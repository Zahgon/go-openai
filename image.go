package openai

import (
	"context"
	"io"
)

const (
	CreateImageSize256x256   = "256x256"
	CreateImageSize512x512   = "512x512"
	CreateImageSize1024x1024 = "1024x1024"

	CreateImageSize1792x1024 = "1792x1024"
	CreateImageSize1024x1792 = "1024x1792"

	CreateImageSize1536x1024 = "1536x1024"
	CreateImageSize1024x1536 = "1024x1536"
)

const (
	CreateImageResponseFormatB64JSON = "b64_json"
	CreateImageResponseFormatURL     = "url"
)

const (
	CreateImageModelDallE2    = "dall-e-2"
	CreateImageModelDallE3    = "dall-e-3"
	CreateImageModelGptImage1 = "gpt-image-1"
)

const (
	CreateImageQualityHD       = "hd"
	CreateImageQualityStandard = "standard"

	CreateImageQualityHigh   = "high"
	CreateImageQualityMedium = "medium"
	CreateImageQualityLow    = "low"
)

const (
	CreateImageStyleVivid   = "vivid"
	CreateImageStyleNatural = "natural"
)

const (
	CreateImageBackgroundTransparent = "transparent"
	CreateImageBackgroundOpaque      = "opaque"
)

const (
	CreateImageModerationLow = "low"
)

const (
	CreateImageOutputFormatPNG  = "png"
	CreateImageOutputFormatJPEG = "jpeg"
	CreateImageOutputFormatWEBP = "webp"
)

type ImageRequest struct {
	Prompt            string `json:"prompt,omitempty"`
	Model             string `json:"model,omitempty"`
	N                 int    `json:"n,omitempty"`
	Quality           string `json:"quality,omitempty"`
	Size              string `json:"size,omitempty"`
	Style             string `json:"style,omitempty"`
	ResponseFormat    string `json:"response_format,omitempty"`
	User              string `json:"user,omitempty"`
	Background        string `json:"background,omitempty"`
	Moderation        string `json:"moderation,omitempty"`
	OutputCompression int    `json:"output_compression,omitempty"`
	OutputFormat      string `json:"output_format,omitempty"`
}

type ImageResponse struct {
	Created int64                    `json:"created,omitempty"`
	Data    []ImageResponseDataInner `json:"data,omitempty"`
	Usage   ImageResponseUsage       `json:"usage,omitempty"`

	httpHeader
}

type ImageResponseInputTokensDetails struct {
	TextTokens  int `json:"text_tokens,omitempty"`
	ImageTokens int `json:"image_tokens,omitempty"`
}

type ImageResponseUsage struct {
	TotalTokens        int                             `json:"total_tokens,omitempty"`
	InputTokens        int                             `json:"input_tokens,omitempty"`
	OutputTokens       int                             `json:"output_tokens,omitempty"`
	InputTokensDetails ImageResponseInputTokensDetails `json:"input_tokens_details,omitempty"`
}

type ImageResponseDataInner struct {
	URL           string `json:"url,omitempty"`
	B64JSON       string `json:"b64_json,omitempty"`
	RevisedPrompt string `json:"revised_prompt,omitempty"`
}

func (c *Client) CreateImage(ctx context.Context, request ImageRequest) (response ImageResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ImageResponse), nil
}

func WrapReader(rdr io.Reader, filename string, contentType string) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

type file struct {
	io.Reader
	name        string
	contentType string
}

func (f file) Name() string { _ = "STUB: not implemented"; return "" }

func (f file) ContentType() string { _ = "STUB: not implemented"; return "" }

type ImageEditRequest struct {
	Image          io.Reader `json:"image,omitempty"`
	Mask           io.Reader `json:"mask,omitempty"`
	Prompt         string    `json:"prompt,omitempty"`
	Model          string    `json:"model,omitempty"`
	N              int       `json:"n,omitempty"`
	Size           string    `json:"size,omitempty"`
	ResponseFormat string    `json:"response_format,omitempty"`
	Quality        string    `json:"quality,omitempty"`
	User           string    `json:"user,omitempty"`
}

func (c *Client) CreateEditImage(ctx context.Context, request ImageEditRequest) (response ImageResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ImageResponse), nil
}

type ImageVariRequest struct {
	Image          io.Reader `json:"image,omitempty"`
	Model          string    `json:"model,omitempty"`
	N              int       `json:"n,omitempty"`
	Size           string    `json:"size,omitempty"`
	ResponseFormat string    `json:"response_format,omitempty"`
	User           string    `json:"user,omitempty"`
}

func (c *Client) CreateVariImage(ctx context.Context, request ImageVariRequest) (response ImageResponse, err error) {
	_ = "STUB: not implemented"
	return *new(ImageResponse), nil
}
