package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

var VERSION string

func parseSlackChannels(channels string) []string {
	list := []string{}

	if strings.Trim(channels, " ") == "" {
		return list
	}

	parts := strings.Split(channels, ENV_SLACK_CHANNEL_DELIMITER)

	for _, i := range parts {
		channel := SlackChannel{Name: strings.Trim(i, " ")}
		if channel.Valid() {
			list = append(list, channel.String())
		}
	}

	return list
}

func Handler(ctx context.Context, event events.CloudWatchEvent) {
  fmt.Printf("App Version: %s\n", VERSION)
	detail := processEventDetail(event)
	fmt.Printf("Event Detail: %v\n", detail)

	slackChannels := os.Getenv(ENV_SLACK_CHANNEL)
	channels := parseSlackChannels(slackChannels)
	if len(channels) == 0 {
		panic("No Slack channels have been configured. Aborting")
	}
	fmt.Printf("Slack channels to notify: %v\n", channels)

	var fmtString string
	switch detail.EventName {
	case EVENT_ADD_USER:
		fmtString = ADD_USER_SLACK_MSG
	case EVENT_REMOVE_USER:
		fmtString = REMOVE_USER_SLACK_MSG
	default:
		fmt.Printf("Unsupported event received: %s\n", detail.EventName)
		return
	}

	msg := SlackMessage{
		Message:   fmt.Sprintf(fmtString, detail.RequestParameters.UserName, detail.RequestParameters.GroupName, detail.SourceIPAddress),
		UserName:  DEFAULT_USERNAME,
		IconEmoji: DEFAULT_EMOJI,
	}

	var url strings.Builder
	url.WriteString("https://hooks.slack.com/services/")
	url.WriteString(os.Getenv(ENV_SLACK_WEBHOOK_URL))
	for _, i := range channels {
		msg.Send(i, url.String())
	}
}

func main() {
	lambda.Start(Handler)
}
