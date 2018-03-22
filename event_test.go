package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestProcessEventDetail(t *testing.T) {
	t.Parallel()

	expected := EventDetail{}
	expected.UserIdentity.SessionContext.SessionIssuer.UserName = "johndoe"
	expected.EventName = "AddUserToGroup"
	expected.RequestParameters.GroupName = "administrators"
	expected.RequestParameters.UserName = "alice"
	expected.SourceIPAddress = "1.2.3.4"

	payload := `{ "eventVersion": "1.02", "userIdentity": { "type": "AssumedRole", "principalId": "AROAIPREDACTED76ERTC:john.doe", "arn": "arn:aws:sts::0123456789012:assumed-role/johndoe/john.doe", "accountId": "0123456789012", "accessKeyId": "ASIAJREDACTEDUOPCWVA", "sessionContext": { "attributes": { "mfaAuthenticated": "true", "creationDate": "2018-03-21T19:57:20Z" }, "sessionIssuer": { "type": "Role", "principalId": "AROAIP2REDACTED76ERTC", "arn": "arn:aws:iam::0123456789012:role/johndoe", "accountId": "0123456789012", "userName": "johndoe" } } }, "eventTime": "2018-03-21T19:59:36Z", "eventSource": "iam.amazonaws.com", "eventName": "AddUserToGroup", "awsRegion": "us-east-1", "sourceIPAddress": "1.2.3.4", "userAgent": "console.amazonaws.com", "requestParameters": { "groupName": "administrators", "userName": "alice" }, "responseElements": null, "requestID": "61726b23-xxxx-yyyy-9f6e-ed128332509c", "eventID": "ba779bda-xxxx-yyyy-8501-34c2e1a5488d", "eventType": "AwsApiCall" }`

	event := events.CloudWatchEvent{
		Detail: []byte(payload),
	}

	actual := processEventDetail(event)

	if actual != expected {
		t.Errorf("[case1] Expected %v but got %v", expected, actual)
	}
}
