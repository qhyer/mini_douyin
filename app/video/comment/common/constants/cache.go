package constants

import (
	"fmt"
	"time"
)

var CommentListCacheKey = func(videoId int64) string {
	return fmt.Sprintf("COMMENT_LIST_%d", videoId)
}

const CommentListCacheExpiration = 3 * time.Minute

var CommentCountCacheKey = func(videoId int64) string {
	return fmt.Sprintf("COMMENT_COUNT_%d", videoId)
}

const CommentCountCacheExpiration = 3 * time.Minute

var CommentInfoCacheKey = func(commentId int64) string {
	return fmt.Sprintf("COMMENT_INFO_%d", commentId)
}

const CommentInfoCacheExpiration = 3 * time.Minute

var CommentStatTempCacheKey = func(procId int) string {
	return fmt.Sprintf("COMMENT_STAT_TEMP_%d", procId)
}
