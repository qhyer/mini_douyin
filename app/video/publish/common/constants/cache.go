package constants

import (
	"fmt"
	"time"
)

const VideoCacheExpiration = 3 * time.Minute

var VideoCacheKey = func(vid int64) string {
	return fmt.Sprintf("VIDEO_INFO_%d", vid)
}

var UserPublishedVidListCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_PUB_VID_LIST_%d", uid)
}

var UserPublishedVidCountCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_PUB_VID_COUNT_%d", uid)
}
