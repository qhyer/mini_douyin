package constants

import (
	"fmt"
	"time"
)

var UserFavoriteListCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FAV_LIST:%d", userId)
}

const UserFavoriteListCacheExpiration = 3 * time.Minute

var UserVideoFavoriteCountCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FAV_COUNT:%d", userId)
}

const UserVideoFavoriteCountCacheExpiration = 3 * time.Minute

var UserVideoFavoritedCountCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FAVD_COUNT:%d", userId)
}

const UserVideoFavoritedCountCacheExpiration = 3 * time.Minute

var VideoFavoritedCountCacheKey = func(videoId int64) string {
	return fmt.Sprintf("VIDEO_FAVD_COUNT:%d", videoId)
}

const VideoFavoritedCountCacheExpiration = 3 * time.Minute
