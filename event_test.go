package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestProcessEventDetail(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		Payload     string
		Expected    EventDetail
	}{
		{
			"case1",
			`{
				"eventVersion": "1.02",
				"userIdentity": {
					"type": "AssumedRole",
					"principalId": "AROAIPREDACTED76ERTC:john.doe",
					"arn": "arn:aws:sts::0123456789012:assumed-role/johndoe/john.doe",
					"accountId": "0123456789012",
					"accessKeyId": "ASIAJREDACTEDUOPCWVA",
					"sessionContext": {
						"attributes": {
							"mfaAuthenticated": "true",
							"creationDate": "2018-03-21T19:57:20Z"
						},
						"sessionIssuer": {
							"type": "Role",
							"principalId": "AROAIP2REDACTED76ERTC",
							"arn": "arn:aws:iam::0123456789012:role/johndoe",
							"accountId": "0123456789012",
							"userName": "johndoe"
						}
					}
				},
				"eventTime": "2018-03-21T19:59:36Z",
				"eventSource": "iam.amazonaws.com",
				"eventName": "AddUserToGroup",
				"awsRegion": "us-east-1",
				"sourceIPAddress": "1.2.3.4",
				"userAgent": "console.amazonaws.com",
				"requestParameters": {
					"groupName": "administrators",
					"userName": "alice"
				},
				"responseElements": null,
				"requestID": "61726b23-xxxx-yyyy-9f6e-ed128332509c",
				"eventID": "ba779bda-xxxx-yyyy-8501-34c2e1a5488d",
				"eventType": "AwsApiCall"
			}`,
			EventDetail{
				UserIdentity: UserIdentity{
					SessionContext: SessionContext{
						SessionIssuer: SessionIssuer{
							UserName: "johndoe",
						},
				 	},
					Arn: "arn:aws:sts::0123456789012:assumed-role/johndoe/john.doe",
				},
				EventName: "AddUserToGroup",
				RequestParameters: RequestParameters{
					GroupName: "administrators",
					UserName: "alice",
				},
				SourceIPAddress: "1.2.3.4",
			},
		},
		{
			"case2",
			`{
				"eventVersion": "1.05",
				"userIdentity": {
						"type": "IAMUser",
						"principalId": "AIDAIEAS7DSJHPLYZGRT6",
						"arn": "arn:aws:iam::002682819933:user/warren.wegner",
						"accountId": "002682819933",
						"accessKeyId": "ASIAQBH7ISVOVZVTWMUL",
						"userName": "warren.wegner",
						"sessionContext": {
								"sessionIssuer": {},
								"webIdFederationData": {},
								"attributes": {
										"mfaAuthenticated": "true",
										"creationDate": "2020-10-23T16:16:23Z"
								}
						}
				},
				"eventTime": "2020-10-23T16:20:43Z",
				"eventSource": "iam.amazonaws.com",
				"eventName": "RemoveUserFromGroup",
				"awsRegion": "us-east-1",
				"sourceIPAddress": "1.2.3.4",
				"userAgent": "console.amazonaws.com",
				"errorCode": "AccessDenied",
				"errorMessage": "User: arn:aws:iam::002682819933:user/warren.wegner is not authorized to perform: iam:RemoveUserFromGroup on resource: group iam-group-content-tribe-Group-VCVVSEI39MNZ",
				"requestParameters": null,
				"responseElements": null,
				"requestID": "66b018a3-d552-4fbf-bfdc-347cee09b5ce",
				"eventID": "245bfdbd-5576-435d-8ec6-a060003f27c5",
				"eventType": "AwsApiCall",
				"recipientAccountId": "002682819933"
			}`,
			EventDetail{
				ErrorMessage: "User: arn:aws:iam::002682819933:user/warren.wegner is not authorized to perform: iam:RemoveUserFromGroup on resource: group iam-group-content-tribe-Group-VCVVSEI39MNZ",
				ErrorCode: "AccessDenied",
				UserIdentity: UserIdentity{
					SessionContext: SessionContext{
						SessionIssuer: SessionIssuer{
							UserName: "",
						},
				 	},
					Arn: "arn:aws:iam::002682819933:user/warren.wegner",
				},
				EventName: "RemoveUserFromGroup",
				RequestParameters: RequestParameters{
					GroupName: "",
					UserName: "",
				},
				SourceIPAddress: "1.2.3.4",
			},
		},
	}

	for _, tc := range cases {
		event := events.CloudWatchEvent{
			Detail: []byte(tc.Payload),
		}

		actual := processEventDetail(event)

		if actual != tc.Expected {
			t.Errorf("[%s] Expected %v but got %v", tc.Name, tc.Expected, actual)
		}
	}
}
