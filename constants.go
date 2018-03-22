package main

const (
	ENV_SLACK_CHANNEL           string = "SLACK_CHANNELS"
	ENV_SLACK_CHANNEL_DELIMITER string = ","
	ENV_SLACK_WEBHOOK_URL       string = "SLACK_WEBHOOK_URL"

	EVENT_ADD_USER    string = "AddUserToGroup"
	EVENT_REMOVE_USER string = "RemoveUserFromGroup"

	ADD_USER_SLACK_MSG    string = ":red-light: User %s was added to group %s (IP: %s)"
	REMOVE_USER_SLACK_MSG string = ":green-light: User %s was removed from group %s (IP: %s)"

	DEFAULT_EMOJI    string = ":aws:"
	DEFAULT_USERNAME string = "Escalated Privileges Watcher"

	SLACK_CHANNEL_PREFIX string = "#"
	SLACK_USER_PREFIX    string = "@"
)
