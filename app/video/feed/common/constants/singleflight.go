package constants

import "fmt"

var SFUserFavoriteVideoIdListKey = func(userId int64) string {
	return fmt.Sprintf("user_favorite_video_%d", userId)
}
