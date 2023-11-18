package constants

import "fmt"

const relationRecordSharding = 100

var RelationRecordTable = func(userId int64) string {
	return fmt.Sprintf("relation_record_%d", userId%relationRecordSharding)
}

const relationCountSharding = 100

var RelationCountTable = func(userId int64) string {
	return fmt.Sprintf("relation_count_%d", userId%relationCountSharding)
}
