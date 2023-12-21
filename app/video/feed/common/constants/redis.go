package constants

import (
	"fmt"
	"time"
)

var VideoInfoCacheKey = func(videoID int64) string {
	return fmt.Sprintf("video:info:%d", videoID)
}

const VideoInfoCacheExpiration = 3 * time.Minute

const LeastVideoFavoritedCount = 100000
