package constants

import (
	"fmt"
	"time"
)

var UserFavoriteListCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FAV_LIST:%d", userId)
}

const UserFavoriteListCacheExpiration = 3 * time.Minute

var UserFavoriteCountCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FAV_COUNT:%d", userId)
}

const UserFavoriteCountCacheExpiration = 3 * time.Minute

var UserFavoritedCountCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FAVD_COUNT:%d", userId)
}

const UserFavoritedCountCacheExpiration = 3 * time.Minute

var VideoFavoritedCountCacheKey = func(videoId int64) string {
	return fmt.Sprintf("VIDEO_FAVD_COUNT:%d", videoId)
}

const VideoFavoritedCountCacheExpiration = 3 * time.Minute

var UserFavoriteStatTempCacheKey = func(procId int) string {
	return fmt.Sprintf("USER_FAV_STAT_TEMP:%d", procId)
}

var UserFavoritedStatTempCacheKey = func(procId int) string {
	return fmt.Sprintf("USER_FAVD_STAT_TEMP:%d", procId)
}

var VideoFavoritedStatTempCacheKey = func(procId int) string {
	return fmt.Sprintf("VIDEO_FAVD_STAT_TEMP:%d", procId)
}
