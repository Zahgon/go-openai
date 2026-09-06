package openai

import (
	"net/http"
	"time"
)

type RateLimitHeaders struct {
	LimitRequests     int       `json:"x-ratelimit-limit-requests"`
	LimitTokens       int       `json:"x-ratelimit-limit-tokens"`
	RemainingRequests int       `json:"x-ratelimit-remaining-requests"`
	RemainingTokens   int       `json:"x-ratelimit-remaining-tokens"`
	ResetRequests     ResetTime `json:"x-ratelimit-reset-requests"`
	ResetTokens       ResetTime `json:"x-ratelimit-reset-tokens"`
}

type ResetTime string

func (r ResetTime) String() string { _ = "STUB: not implemented"; return "" }

func (r ResetTime) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func newRateLimitHeaders(h http.Header) RateLimitHeaders {
	_ = "STUB: not implemented"
	return *new(RateLimitHeaders)
}
