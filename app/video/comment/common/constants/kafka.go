package constants

import "fmt"

const CommentActionTopic string = "comment_action"

var CommentActionKafkaKey = func(userId int64) string {
	return fmt.Sprintf("%d", userId)
}

const UpdateCommentCountTopic string = "update_comment_count"

var UpdateCommentCountKafkaKey = func(videoId int64) string {
	return fmt.Sprintf("%d", videoId)
}
