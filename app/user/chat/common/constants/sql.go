package constants

import "fmt"

const messageRecordSharding = 100

var MessageRecordTable = func(userId, toUserId int64) string {
	return fmt.Sprintf("message_record_%d", (userId+toUserId)%messageRecordSharding)
}
