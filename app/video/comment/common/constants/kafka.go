package constants

import "fmt"

const CommentActionTopic string = "comment_action"

var CommentActionKafkaKey = func(userId int64) string {
	return fmt.Sprintf("%d", userId)
}
