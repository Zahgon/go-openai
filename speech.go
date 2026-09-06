package openai

import (
	"context"
)

type SpeechModel string

const (
	TTSModel1         SpeechModel = "tts-1"
	TTSModel1HD       SpeechModel = "tts-1-hd"
	TTSModelCanary    SpeechModel = "canary-tts"
	TTSModelGPT4oMini SpeechModel = "gpt-4o-mini-tts"
)

type SpeechVoice string

const (
	VoiceAlloy   SpeechVoice = "alloy"
	VoiceAsh     SpeechVoice = "ash"
	VoiceBallad  SpeechVoice = "ballad"
	VoiceCoral   SpeechVoice = "coral"
	VoiceEcho    SpeechVoice = "echo"
	VoiceFable   SpeechVoice = "fable"
	VoiceOnyx    SpeechVoice = "onyx"
	VoiceNova    SpeechVoice = "nova"
	VoiceShimmer SpeechVoice = "shimmer"
	VoiceVerse   SpeechVoice = "verse"
)

type SpeechResponseFormat string

const (
	SpeechResponseFormatMp3  SpeechResponseFormat = "mp3"
	SpeechResponseFormatOpus SpeechResponseFormat = "opus"
	SpeechResponseFormatAac  SpeechResponseFormat = "aac"
	SpeechResponseFormatFlac SpeechResponseFormat = "flac"
	SpeechResponseFormatWav  SpeechResponseFormat = "wav"
	SpeechResponseFormatPcm  SpeechResponseFormat = "pcm"
)

type CreateSpeechRequest struct {
	Model          SpeechModel          `json:"model"`
	Input          string               `json:"input"`
	Voice          SpeechVoice          `json:"voice"`
	Instructions   string               `json:"instructions,omitempty"`
	ResponseFormat SpeechResponseFormat `json:"response_format,omitempty"`
	Speed          float64              `json:"speed,omitempty"`
}

func (c *Client) CreateSpeech(ctx context.Context, request CreateSpeechRequest) (response RawResponse, err error) {
	_ = "STUB: not implemented"
	return *new(RawResponse), nil
}
