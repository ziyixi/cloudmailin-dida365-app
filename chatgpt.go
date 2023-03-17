package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func summaryByChatGPT(content string) string {
	client := openai.NewClient(os.Getenv("openai_api_key"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: `Summary the following email as markdown bullet points concisely.			

					` + content,
				},
			},
		},
	)

	if err != nil {
		return fmt.Sprintf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}
