package constants

import "fmt"

const relationRecordSharding = 100

var FollowRecordTable = func(userId int64) string {
	return fmt.Sprintf("follow_record_%d", userId%relationRecordSharding)
}

var FollowerRecordTable = func(userId int64) string {
	return fmt.Sprintf("follower_record_%d", userId%relationRecordSharding)
}

const relationCountSharding = 100

var RelationCountTable = func(userId int64) string {
	return fmt.Sprintf("relation_count_%d", userId%relationCountSharding)
}
