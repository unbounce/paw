package main

const (
	ENV_SLACK_CHANNEL           string = "SLACK_CHANNELS"
	ENV_SLACK_CHANNEL_DELIMITER string = ","
	ENV_SLACK_WEBHOOK_URL       string = "SLACK_WEBHOOK_URL"

	EVENT_ADD_USER    string = "AddUserToGroup"
	EVENT_REMOVE_USER string = "RemoveUserFromGroup"
	GROUP_BASE_URL  string = "https://console.aws.amazon.com/iam/home?region=us-east-1#/groups/"
	ROLE_BASE_URL   string = "https://console.aws.amazon.com/iam/home?region=us-east-1#/roles/"
	USER_BASE_URL   string = "https://console.aws.amazon.com/iam/home?region=us-east-1#/users/"

	ADD_USER_SLACK_MSG    string = ":red-light: User %s was added to group %s (*Arn:* %s *IP:* %s)"
	REMOVE_USER_SLACK_MSG string = ":green-light: User %s was removed from group %s (*Arn:* %s *IP:* %s)"
	ERROR_SLACK_MSG string = ":x: IAM update failure: %s (*IP:* %s)"

	DEFAULT_EMOJI    string = ":aws:"
	DEFAULT_USERNAME string = "Escalated Privileges Watcher"

	SLACK_CHANNEL_PREFIX string = "#"
	SLACK_USER_PREFIX    string = "@"
)
