package constants

import (
	"fmt"
	"time"
)

var ChatConversationCacheKey = func(uid1, uid2 int64) string {
	if uid1 < uid2 {
		return fmt.Sprintf("chat:conversation:%d:%d", uid1, uid2)
	} else {
		return fmt.Sprintf("chat:conversation:%d:%d", uid2, uid1)
	}
}

const ChatLatestMsgCacheExpiration = 3 * time.Minute

var ChatConversationLatestMsgCacheKey = func(uid1, uid2 int64) string {
	if uid1 < uid2 {
		return fmt.Sprintf("chat:conversation:latestMsg:%d:%d", uid1, uid2)
	} else {
		return fmt.Sprintf("chat:conversation:latestMsg:%d:%d", uid2, uid1)
	}
}
