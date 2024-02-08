package constants

import "fmt"

const (
	FavoriteVideoRecordSharding    int64 = 100
	UserFavoriteVideoCountSharding int64 = 100
	VideoFavoritedCountSharding    int64 = 100
)

var FavoriteVideoRecordTableName = func(userId int64) string {
	return fmt.Sprintf("favorite_video_record_%d", userId%FavoriteVideoRecordSharding)
}

var UserFavoriteVideoCountTableName = func(userId int64) string {
	return fmt.Sprintf("user_favorite_video_count_%d", userId%UserFavoriteVideoCountSharding)
}

var VideoFavoritedCountTableName = func(videoId int64) string {
	return fmt.Sprintf("video_favorited_count_%d", videoId%VideoFavoritedCountSharding)
}
