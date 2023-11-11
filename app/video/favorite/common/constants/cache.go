package constants

import (
	"fmt"
	"time"
)

var UserFavoriteListCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_FAV_LIST:%d", uid)
}

const UserFavoriteListCacheExpiration = 3 * time.Minute
