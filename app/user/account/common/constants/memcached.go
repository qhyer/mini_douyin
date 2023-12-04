package constants

import (
	"fmt"
	"time"
)

var UserFollowListCacheKey = func(userId int64) string {
	return fmt.Sprintf("user_follow_list_%d", userId)
}

const UserFollowListCacheExpiration = 3 * time.Minute

var UserFollowerListCacheKey = func(userId int64) string {
	return fmt.Sprintf("user_follower_list_%d", userId)
}

const UserFollowerListCacheExpiration = 3 * time.Minute

var UserFriendListCacheKey = func(userId int64) string {
	return fmt.Sprintf("user_friend_list_%d", userId)
}

const UserFriendListCacheExpiration = 3 * time.Minute
