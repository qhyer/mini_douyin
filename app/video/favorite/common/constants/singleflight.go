package constants

import "fmt"

var SFUserFavoriteIdListKey = func(userId int64) string {
	return fmt.Sprintf("sf_user_fav_id_list_%d", userId)
}

var SFCountUserFavoriteKey = func(userId int64) string {
	return fmt.Sprintf("sf_count_user_favorite_%d", userId)
}

var SFCountUserFavoritedKey = func(userId int64) string {
	return fmt.Sprintf("sf_count_user_favorited_%d", userId)
}

var SFCountVideoFavoritedKey = func(videoId int64) string {
	return fmt.Sprintf("sf_count_video_favorited_%d", videoId)
}
