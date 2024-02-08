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

var UserFriendListCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FRIEND_LIST:%d", userId)
}

const UserFriendListCacheExpiration = 3 * time.Minute

var UserFollowListCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FOLLOW_LIST:%d", userId)
}

const UserFollowListCacheExpiration = 3 * time.Minute

var UserFollowerListCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FOLLOWER_LIST:%d", userId)
}

const UserFollowerListCacheExpiration = 3 * time.Minute

var UserFollowBloomCacheKey = func(userId int64) string {
	return fmt.Sprintf("USER_FOLLOW_BLOOM:%d", userId)
}

var UserFollowTempCountKey = func(procId int) string {
	return fmt.Sprintf("USER_FOLLOW_TEMP_COUNT:%d", procId)
}

var UserFollowerTempCountKey = func(procId int) string {
	return fmt.Sprintf("USER_FOLLOWER_TEMP_COUNT:%d", procId)
}
