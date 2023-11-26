package constants

import "fmt"

var SFUserPublishedVideoKey = func(userId int64) string {
	return fmt.Sprintf("sf_user_published_video_%d", userId)
}

var SFLatestPublishedVideoKey = func(ts int64, limit int) string {
	return fmt.Sprintf("sf_latest_published_video_%d_%d", ts, limit)
}

var SFVideoKey = func(videoId int64) string {
	return fmt.Sprintf("sf_video_%d", videoId)
}

var SFVideoCountKey = func(userId int64) string {
	return fmt.Sprintf("sf_video_count_%d", userId)
}
