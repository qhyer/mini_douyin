package constants

import "fmt"

const CommentRecordSharding = 100

var CommentRecordTableName = func(videoId int64) string {
	return fmt.Sprintf("comment_record_%d", videoId%CommentRecordSharding)
}

const CommentCountSharding = 100

var CommentCountTableName = func(videoId int64) string {
	return fmt.Sprintf("comment_count_%d", videoId%CommentCountSharding)
}

const CommentQueryLimit = 100
