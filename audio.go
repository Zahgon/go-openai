package openai

import (
	"context"
	"io"

	utils "github.com/sashabaranov/go-openai/internal"
)

const (
	Whisper1               = "whisper-1"
	GPT4oTranscribe        = "gpt-4o-transcribe"
	GPT4oMiniTranscribe    = "gpt-4o-mini-transcribe"
	GPT4oTranscribeDiarize = "gpt-4o-transcribe-diarize"
	GPTTranscribe          = "gpt-transcribe"
)

type AudioResponseFormat string

const (
	AudioResponseFormatJSON        AudioResponseFormat = "json"
	AudioResponseFormatText        AudioResponseFormat = "text"
	AudioResponseFormatSRT         AudioResponseFormat = "srt"
	AudioResponseFormatVerboseJSON AudioResponseFormat = "verbose_json"
	AudioResponseFormatVTT         AudioResponseFormat = "vtt"
)

type TranscriptionTimestampGranularity string

const (
	TranscriptionTimestampGranularityWord    TranscriptionTimestampGranularity = "word"
	TranscriptionTimestampGranularitySegment TranscriptionTimestampGranularity = "segment"
)

type TranscriptionChunkingStrategy struct {
	Type              string  `json:"type"`
	PrefixPaddingMs   int     `json:"prefix_padding_ms,omitempty"`
	SilenceDurationMs int     `json:"silence_duration_ms,omitempty"`
	Threshold         float64 `json:"threshold,omitempty"`
}

type AudioRequest struct {
	Model string

	FilePath string

	Reader io.Reader

	Prompt                 string
	Temperature            float32
	Language               string
	Format                 AudioResponseFormat
	TimestampGranularities []TranscriptionTimestampGranularity

	ChunkingStrategy any
}

type AudioResponse struct {
	Task     string  `json:"task"`
	Language string  `json:"language"`
	Duration float64 `json:"duration"`
	Segments []struct {
		ID               int     `json:"id"`
		Seek             int     `json:"seek"`
		Start            float64 `json:"start"`
		End              float64 `json:"end"`
		Text             string  `json:"text"`
		Tokens           []int   `json:"tokens"`
		Temperature      float64 `json:"temperature"`
		AvgLogprob       float64 `json:"avg_logprob"`
		CompressionRatio float64 `json:"compression_ratio"`
		NoSpeechProb     float64 `json:"no_speech_prob"`
		Transient        bool    `json:"transient"`
	} `json:"segments"`
	Words []struct {
		Word  string  `json:"word"`
		Start float64 `json:"start"`
		End   float64 `json:"end"`
	} `json:"words"`
	Text string `json:"text"`

	httpHeader
}

type audioTextResponse struct {
	Text string `json:"text"`

	httpHeader
}

func (r *audioTextResponse) ToAudioResponse() AudioResponse {
	_ = "STUB: not implemented"
	return *new(AudioResponse)
}

func (c *Client) CreateTranscription(
	ctx context.Context,
	request AudioRequest,
) (response AudioResponse, err error) {
	_ = "STUB: not implemented"
	return *new(AudioResponse), nil
}

func (c *Client) CreateTranslation(
	ctx context.Context,
	request AudioRequest,
) (response AudioResponse, err error) {
	_ = "STUB: not implemented"
	return *new(AudioResponse), nil
}

func (c *Client) callAudioAPI(
	ctx context.Context,
	request AudioRequest,
	endpointSuffix string,
) (response AudioResponse, err error) {
	_ = "STUB: not implemented"
	return *new(AudioResponse), nil
}

func (r AudioRequest) HasJSONResponse() bool { _ = "STUB: not implemented"; return false }

func audioMultipartForm(request AudioRequest, b utils.FormBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTimestampGranularities(granularities []TranscriptionTimestampGranularity, b utils.FormBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func writeChunkingStrategy(strategy any, b utils.FormBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func createFileField(request AudioRequest, b utils.FormBuilder) error {
	_ = "STUB: not implemented"
	return nil
}
