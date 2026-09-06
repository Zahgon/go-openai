package openai

import (
	"net/http"
)

const (
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300

	azureAPIPrefix         = "openai"
	azureDeploymentsPrefix = "deployments"

	AnthropicAPIVersion = "2023-06-01"
)

type APIType string

const (
	APITypeOpenAI          APIType = "OPEN_AI"
	APITypeAzure           APIType = "AZURE"
	APITypeAzureAD         APIType = "AZURE_AD"
	APITypeCloudflareAzure APIType = "CLOUDFLARE_AZURE"
	APITypeAnthropic       APIType = "ANTHROPIC"
)

const AzureAPIKeyHeader = "api-key"

const defaultAssistantVersion = "v2"

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientConfig struct {
	authToken string

	BaseURL              string
	OrgID                string
	APIType              APIType
	APIVersion           string
	AssistantVersion     string
	AzureModelMapperFunc func(model string) string
	HTTPClient           HTTPDoer

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	_ = "STUB: not implemented"
	return *new(ClientConfig)
}

func DefaultAzureConfig(apiKey, baseURL string) ClientConfig {
	_ = "STUB: not implemented"
	return *new(ClientConfig)
}

func DefaultAnthropicConfig(apiKey, baseURL string) ClientConfig {
	_ = "STUB: not implemented"
	return *new(ClientConfig)
}

func (ClientConfig) String() string { _ = "STUB: not implemented"; return "" }

func (c ClientConfig) GetAzureDeploymentByModel(model string) string {
	_ = "STUB: not implemented"
	return ""
}
