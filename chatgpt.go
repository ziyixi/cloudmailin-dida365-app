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
					Content: `Summarize the following email in a way that highlights the main idea using bullet points in Markdown format so I can easily recall it later, it can use sub-bullet points to be more organized, and it shouldn't include level 1 Headings:			

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
