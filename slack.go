package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const (
	MSG_WAIT int = 1 // seconds
)

type SlackMessage struct {
	Message   string `json:"text"`
	Channel   string `json:"channel"`
	UserName  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

type SlackChannel struct {
	Name string
}

func (c SlackChannel) Valid() bool {
	return c.Name != ""
}

func (c SlackChannel) isRoutable() bool {
	return strings.HasPrefix(c.Name, SLACK_CHANNEL_PREFIX) || strings.HasPrefix(c.Name, SLACK_USER_PREFIX)
}

func (c *SlackChannel) Fix() {
	var fix strings.Builder
	fix.WriteString(SLACK_CHANNEL_PREFIX) // assume they meant a channel
	fix.WriteString(c.Name)
	c.Name = fix.String()
}

func (c *SlackChannel) String() string {
	if !c.isRoutable() {
		c.Fix()
		return c.Name
	}
	return c.Name
}

func (m *SlackMessage) Send(channel string, url string) {
	m.Channel = channel

	payload, err := json.Marshal(m)
	if err != nil {
		panic(err.Error())
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(payload))
	time.Sleep(time.Duration(MSG_WAIT) * time.Second)
}
