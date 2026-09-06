package openai

import (
	"errors"
)

var (
	ErrO1MaxTokensDeprecated                   = errors.New("this model is not supported MaxTokens, please use MaxCompletionTokens")                               //nolint:lll
	ErrCompletionUnsupportedModel              = errors.New("this model is not supported with this method, please use CreateChatCompletion client method instead") //nolint:lll
	ErrCompletionStreamNotSupported            = errors.New("streaming is not supported with this method, please use CreateCompletionStream")                      //nolint:lll
	ErrCompletionRequestPromptTypeNotSupported = errors.New("the type of CompletionRequest.Prompt only supports string and []string")                              //nolint:lll
)

var (
	ErrO1BetaLimitationsMessageTypes = errors.New("this model has beta-limitations, user and assistant messages only, system messages are not supported")       //nolint:lll
	ErrO1BetaLimitationsTools        = errors.New("this model has beta-limitations, tools, function calling, and response format parameters are not supported") //nolint:lll

	ErrO1BetaLimitationsLogprobs = errors.New("this model has beta-limitations, logprobs not supported")                                                                               //nolint:lll
	ErrO1BetaLimitationsOther    = errors.New("this model has beta-limitations, temperature, top_p and n are fixed at 1, while presence_penalty and frequency_penalty are fixed at 0") //nolint:lll
)

var (
	//nolint:lll
	ErrReasoningModelMaxTokensDeprecated = errors.New("this model is not supported MaxTokens, please use MaxCompletionTokens")
	ErrReasoningModelLimitationsLogprobs = errors.New("this model has beta-limitations, logprobs not supported")                                                                               //nolint:lll
	ErrReasoningModelLimitationsOther    = errors.New("this model has beta-limitations, temperature, top_p and n are fixed at 1, while presence_penalty and frequency_penalty are fixed at 0") //nolint:lll
)

type ReasoningValidator struct{}

func NewReasoningValidator() *ReasoningValidator { _ = "STUB: not implemented"; return nil }

func (v *ReasoningValidator) Validate(request ChatCompletionRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *ReasoningValidator) validateReasoningModelParams(request ChatCompletionRequest) error {
	_ = "STUB: not implemented"
	return nil
}
