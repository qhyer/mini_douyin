package constants

import "fmt"

var SFUserInfoKey = func(userId int64) string {
	return fmt.Sprintf("sf_user_info_%v", userId)
}
