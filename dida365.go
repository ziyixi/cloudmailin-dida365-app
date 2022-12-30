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
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("no .env file found")
	}
	c, err := ticktick.NewClient(os.Getenv("username"), os.Getenv("password"), "dida365")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func CreateDidaTask(c *ticktick.Client, p *parsedPost) (*ticktick.TaskItem, error) {
	// task content
	template := "**From: %v** \n**Date: %v** \n**Received: %v** \n**Subject: %v** \n\n" + strings.Repeat("=", 20) + "\n%v"
	content := fmt.Sprintf(template, p.From, p.Date, p.To, p.Subject, p.Content)
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
