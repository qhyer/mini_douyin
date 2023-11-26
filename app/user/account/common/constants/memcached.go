package constants

import "fmt"

var UserFollowListCacheKey = func(userId int64) string {
	return fmt.Sprintf("user_follow_list_%d", userId)
}

var UserFollowerListCacheKey = func(userId int64) string {
	return fmt.Sprintf("user_follower_list_%d", userId)
}

var UserFriendListCacheKey = func(userId int64) string {
	return fmt.Sprintf("user_friend_list_%d", userId)
}
