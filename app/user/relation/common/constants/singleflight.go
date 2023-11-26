package constants

import "fmt"

var SFFollowListKey = func(userId int64) string {
	return fmt.Sprintf("sf_follow_list_%d", userId)
}

var SFFollowerListKey = func(userId int64) string {
	return fmt.Sprintf("sf_follower_list_%d", userId)
}

var SFFollowCountKey = func(userId int64) string {
	return fmt.Sprintf("sf_follow_count_%d", userId)
}

var SFFollowerCountKey = func(userId int64) string {
	return fmt.Sprintf("sf_follower_count_%d", userId)
}

var SFFriendListKey = func(userId int64) string {
	return fmt.Sprintf("sf_friend_list_%d", userId)
}

var SFIsFollowKey = func(userId, toUserId int64) string {
	return fmt.Sprintf("sf_is_follow_%d_%d", userId, toUserId)
}
