package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/ziyixi/go-ticktick"
)

func CreateDidaTask(p *parsedPost) error {
	// create client
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("no .env file found")
	}
	c, err := ticktick.NewClient(os.Getenv("username"), os.Getenv("password"), "dida365")
	if err != nil {
		return err
	}

	// task content
	template := "**From: %v** \n**Date: %v** \n**Subject: %v** \n\n" + strings.Repeat("=", 20) + "\n%v"
	content := fmt.Sprintf(template, p.From, p.Date, p.Subject, p.Content)
	title := fmt.Sprintf("%v [%v]", p.Subject, p.From)

	// create task
	task, err := ticktick.NewTask(c, title, content, time.Time{}, "")
	if err != nil {
		return err
	}
	_, err = c.CreateTask(task)
	return err
}
