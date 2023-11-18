package constants

import "fmt"

const FavoriteVideoActionTopic string = "favorite_video"

var FavoriteActionKafkaKey = func(userId int64) string {
	return fmt.Sprintf("%d", userId)
}
