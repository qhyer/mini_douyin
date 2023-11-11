package constants

import (
	"fmt"
	"time"
)

var UserCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_%d", uid)
}

const UserCacheExpiration = 3 * time.Minute
