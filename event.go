package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type UserIdentity struct {
	SessionContext SessionContext `json:"sessionContext"`
	Arn            string         `json:"arn"`
}

type SessionContext struct {
	SessionIssuer SessionIssuer `json:"sessionIssuer"`
}

type SessionIssuer struct {
	UserName string `json:"userName"`
}

type RequestParameters struct {
	GroupName string `json:"groupName"`
	UserName  string `json:"userName"`
}

type EventDetail struct {
	UserIdentity      UserIdentity      `json:"userIdentity"`
	EventName         string            `json:"eventName"`
	RequestParameters RequestParameters `json:"requestParameters"`
	SourceIPAddress   string            `json:"sourceIPAddress"`
	ErrorMessage      string            `json:"errorMessage"`
	ErrorCode         string            `json:"errorCode"`
}

func processEventDetail(event events.CloudWatchEvent) EventDetail {
	var data EventDetail

	err := json.Unmarshal(event.Detail, &data)
	if err != nil {
		panic(err.Error())
	}

	return data
}
