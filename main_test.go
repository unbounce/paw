package main

import (
	"testing"
	"fmt"
)

func TestParseSlackChannels(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name        string
		ChannelList string
		Expected    []string
	}{
		{"case1", "a", []string{"#a"}},
		{"case2", "@a", []string{"@a"}},
		{"case3", "#a", []string{"#a"}},
		{"case4", "a,b", []string{"#a", "#b"}},
		{"case5", "#a,b", []string{"#a", "#b"}},
		{"case6", "#a,b", []string{"#a", "#b"}},
		{"case7", "a,#b", []string{"#a", "#b"}},
		{"case8", "#a,#b", []string{"#a", "#b"}},
		{"case9", "", []string{}},
		{"case10", "           ", []string{}},
		{"case11", " ,   ,   , ,, ", []string{}},
		{"case12", " , c ,   ,a,, ", []string{"#c", "#a"}},
		{"case13", "a,#b,@c", []string{"#a", "#b", "@c"}},
	}

	for _, tc := range cases {
		actual := parseSlackChannels(tc.ChannelList)

		if !stringSliceEq(actual, tc.Expected) {
			t.Errorf("[%s] Expected %v but got %v", tc.Name, tc.Expected, actual)
		}
	}
}

func stringSliceEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCreateErrorMessage(t *testing.T) {
	detail := EventDetail{
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
	}

	actual := createErrorMessage(detail)

	expected := SlackMessage{
		Message: fmt.Sprintf(ERROR_SLACK_MSG, "User: arn:aws:iam::002682819933:user/warren.wegner is not authorized to perform: iam:RemoveUserFromGroup on resource: group iam-group-content-tribe-Group-VCVVSEI39MNZ", "1.2.3.4"),
		UserName: DEFAULT_USERNAME,
		IconEmoji: DEFAULT_EMOJI,
	}

	if (actual != expected) {
		t.Errorf("[CreateErrorMessage] Expected %v but go t%v", expected, actual)
	}

}

func TestCreateNotifyMessage(t *testing.T) {
	detail := EventDetail{
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
	}

	actual := createNotifyMessage(detail)

	expected := SlackMessage{
		Message: fmt.Sprintf(ADD_USER_SLACK_MSG, "<https://console.aws.amazon.com/iam/home?region=us-east-1#/users/alice|alice>", "<https://console.aws.amazon.com/iam/home?region=us-east-1#/groups/administrators|administrators>", "arn:aws:sts::0123456789012:assumed-role/johndoe/john.doe", "1.2.3.4"),
		UserName: DEFAULT_USERNAME,
		IconEmoji: DEFAULT_EMOJI,
	}

	if (actual != expected) {
		t.Errorf("[CreateNotifyMessage] Expected %v but go t%v", expected, actual)
	}
}