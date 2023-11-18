package constants

import (
	"fmt"
	"time"
)

var UserFollowCountCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FOLLOW_COUNT:%d", userId)
}

const UserFollowCountCacheExpiration = 3 * time.Minute

var UserFollowerCountCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FOLLOWER_COUNT:%d", userId)
}

const UserFollowerCountCacheExpiration = 3 * time.Minute
