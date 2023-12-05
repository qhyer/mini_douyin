package constants

import "fmt"

const RelationActionTopic = "relation_action"

var RelationActionKafkaKey = func(fromUserId int64) string {
	return fmt.Sprintf("%d", fromUserId)
}

const UpdateUserFollowCountTopic = "update_user_follow_count"

var UpdateUserFollowCountKafkaKey = func(userId int64) string {
	return fmt.Sprintf("%d", userId)
}

const UpdateUserFollowerCountTopic = "update_user_follower_count"

var UpdateUserFollowerCountKafkaKey = func(userId int64) string {
	return fmt.Sprintf("%d", userId)
}
