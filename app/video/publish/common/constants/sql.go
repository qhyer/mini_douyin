package constants

import "fmt"

const PublishRecordTableName = "publish_record"

const PublishCountSharding = 100

var PublishCountTableName = func(userId int64) string {
	return fmt.Sprintf("publish_count_%d", userId%PublishCountSharding)
}
