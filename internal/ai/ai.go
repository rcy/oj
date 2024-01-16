package ai

import (
	"oj/internal/config"

	openai "github.com/sashabaranov/go-openai"
)

var apiKey = config.MustGetenv("OPENAI_API_KEY")

type AI struct {
	Client *openai.Client
}

func New() *AI {
	return &AI{Client: openai.NewClient(apiKey)}
}
