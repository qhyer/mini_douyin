package constants

import "fmt"

var SFVideoCommentList = func(videoId int64) string {
	return fmt.Sprintf("sf_video_comment_list_%d", videoId)
}

var SFVideoCommentCount = func(videoId int64) string {
	return fmt.Sprintf("sf_video_comment_count_%d", videoId)
}
