package event

import "encoding/json"

type RelationStat struct {
	UserId        int64 `json:"user_id"`
	FollowerDelta int64 `json:"follower_delta"`
	FollowDelta   int64 `json:"follow_delta"`
}

func (e *RelationStat) MarshalJson() ([]byte, error) {
	return json.Marshal(e)
}

func (e *RelationStat) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, e)
}
