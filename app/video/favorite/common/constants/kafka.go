package constants

import "fmt"

const FavoriteVideoActionTopic string = "favorite_video"

var FavoriteActionKafkaKey = func(userId int64) string {
	return fmt.Sprintf("%d", userId)
}

const UpdateVideoFavoritedCountTopic string = "update_video_favorited_count"

var UpdateVideoFavoritedCountKafkaKey = func(videoId int64) string {
	return fmt.Sprintf("%d", videoId)
}
