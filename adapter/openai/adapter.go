package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type Config struct {
	Key   string `koanf:"key"`
	Model string `koanf:"model"`
	Role  string `koanf:"role"`
}
type Adapter struct {
	config Config
	client *openai.Client
}

func New(config Config) *Adapter {
	client := openai.NewClient(config.Key)
	return &Adapter{client: client, config: config}
}

func (a *Adapter) JSON(ctx context.Context, data []string, format string) (string, error) {
	inputs := ""
	for _, d := range data {
		inputs += d + "\n"
	}
	prompt := `You are given a sample input data and I need you to do the same to process and standardize the input data into JSON format.
Sample standard format:` + format + `"note that just output only the raw JSON with given input order and without any prettier and use your information to complete empty values!"
inputs:` + inputs

	resp, err := a.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: a.config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    a.config.Role,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
