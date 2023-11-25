package constants

import (
	"fmt"
	"time"
)

var AccountInfoCacheKey = func(userId int64) string {
	return fmt.Sprintf("ACC_INFO_%d", userId)
}

const AccountInfoCacheExpiration = 3 * time.Minute
