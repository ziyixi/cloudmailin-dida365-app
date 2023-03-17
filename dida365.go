package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/ziyixi/go-ticktick"
)

func LoginDidaClient() (*ticktick.Client, error) {
	godotenv.Load()
	username := os.Getenv("dida365_username")
	password := os.Getenv("dida365_password")
	if username == "" || password == "" {
		return nil, fmt.Errorf("no username or password found")
	}

	c, err := ticktick.NewClient(username, password, "dida365")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func CreateDidaTask(c *ticktick.Client, p *parsedPost) (*ticktick.TaskItem, error) {
	// task content
	template := "**From: %v** \n**Date: %v** \n**Received: %v** \n**Subject: %v** \n\n" + strings.Repeat("=", 20) + "\n%v"
	template_tosummary := "Subject: %v\n\n%v"
	tosummary := fmt.Sprintf(template_tosummary, p.Subject, p.Content)
	chatgptSummary := summaryByChatGPT(tosummary)
	content := fmt.Sprintf(template, p.From, p.Date, p.To, p.Subject, chatgptSummary)
	title := fmt.Sprintf("%v [%v]", p.Subject, p.From)

	// create task
	task, err := ticktick.NewTask(c, title, content, time.Time{}, "")
	task.Tags = append(task.Tags, "email")
	if err != nil {
		return nil, err
	}
	task, err = c.CreateTask(task)
	return task, err
}
