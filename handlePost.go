package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/ziyixi/go-ticktick"
)

type parsedPost struct {
	From    string // headers.from
	To      string // headers.to
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
		From:    gjson.Get(s, "headers.from").String(),
		To:      gjson.Get(s, "headers.to").String(),
		Date:    gjson.Get(s, "headers.date").String(),
		Subject: gjson.Get(s, "headers.subject").String(),
		Content: markdown,
	}

	// outlook may have a prefix FW:
	heloDomain := gjson.Get(s, "envelope.helo_domain").String()
	if strings.Contains(heloDomain, "outlook") && strings.HasPrefix(res.Subject, "FW: ") {
		res.Subject = res.Subject[4:]
	}

	// outlook might foward the email in the forwarding format
	if strings.Contains(res.To, "cloudmailin") {
		// parse the correct To
		re, _ := regexp.Compile(`_+\\r\\nFrom: .*?([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`)
		matches := re.FindStringSubmatch(s)
		if len(matches) < 2 {
			res.To = res.From
			res.From = "sender unknown"
		} else {
			res.To = res.From
			res.From = matches[1]
		}
	}

	return &res
}

func HandleCMIPost(c *gin.Context) {
	// check context
	var client *ticktick.Client
	if clientraw, ok := c.Get("client"); !ok {
		newclient, err := LoginDidaClient()
		if err != nil {
			panic(err)
		}
		client = newclient
		c.Set("client", newclient)
	} else {
		client = clientraw.(*ticktick.Client)
	}

	// handle request
	dataRaw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	data := string(dataRaw)
	parsedRes := parseJson(data)
	if len(parsedRes.From) == 0 {
		return
	}

	task, err := CreateDidaTask(client, parsedRes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task_id": task.Id})
}
