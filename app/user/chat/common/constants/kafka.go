package constants

import "fmt"

const SendMsgTopic string = "send_msg"

var SendMsgKafkaKey = func(myUid, hisUid int64) string {
	if myUid > hisUid {
		return fmt.Sprintf("%d_%d", myUid, hisUid)
	} else {
		return fmt.Sprintf("%d_%d", hisUid, myUid)
	}
}
