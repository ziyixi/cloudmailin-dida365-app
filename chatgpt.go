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
					Content: `I want you to act as a personal email assistant. I will give you the content of an email, and you will provide a summary of that topic in the format of markdown bullet points (very important). Your summary should be concise, covering the most important aspect of the topic. Start your summary with an introductory paragraph that gives an overview of the topic, then bullet points and sub bullet points to make it more organized. It shouldn't include any level of Headings. Below is the content of the email:			

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
