package constants

import "fmt"

const (
	favoriteVideoRecordSharding    int64 = 100
	userFavoriteVideoCountSharding int64 = 100
	videoFavoritedCountSharding    int64 = 100
)

var FavoriteVideoRecordTableName = func(userId int64) string {
	return fmt.Sprintf("favorite_video_record_%d", userId%favoriteVideoRecordSharding)
}

var UserFavoriteVideoCountTableName = func(userId int64) string {
	return fmt.Sprintf("user_favorite_video_count_%d", userId%userFavoriteVideoCountSharding)
}

var VideoFavoritedCountTableName = func(videoId int64) string {
	return fmt.Sprintf("video_favorited_count_%d", videoId%videoFavoritedCountSharding)
}
