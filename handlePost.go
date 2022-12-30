package main

import (
	"io"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type parsedPost struct {
	From    string // envelope.from
	Date    string // headers.date
	Subject string // headers.subject
	Content string // md(html)
}

func parseJson(s string) *parsedPost {
	converter := md.NewConverter("", true, nil)
	html := gjson.Get(s, "html").String()
	markdown, err := converter.ConvertString(html)
	if err != nil || len(markdown) == 0 {
		// use plain text instead
		markdown = gjson.Get(s, "plain").String()
	}

	res := parsedPost{
		From:    gjson.Get(s, "envelope.from").String(),
		Date:    gjson.Get(s, "headers.date").String(),
		Subject: gjson.Get(s, "headers.subject").String(),
		Content: markdown,
	}
	return &res
}

func HandleCMIPost(c *gin.Context) {
	dataRaw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	data := string(dataRaw)
	parsedRes := parseJson(data)
	if len(parsedRes.From) == 0 {
		return
	}

	CreateDidaTask(parsedRes)
}
