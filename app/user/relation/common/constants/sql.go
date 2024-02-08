package constants

import "fmt"

const RelationRecordSharding = 100

var FollowRecordTable = func(userId int64) string {
	return fmt.Sprintf("follow_record_%d", userId%RelationRecordSharding)
}

var FollowerRecordTable = func(userId int64) string {
	return fmt.Sprintf("follower_record_%d", userId%RelationRecordSharding)
}

const RelationCountSharding = 100

var RelationCountTable = func(userId int64) string {
	return fmt.Sprintf("relation_count_%d", userId%RelationCountSharding)
}
