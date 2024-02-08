package event

import "encoding/json"

type CommentStat struct {
	VideoId int64 `json:"video_id"`
	Delta   int64 `json:"delta"`
}

func (c *CommentStat) MarshalJson() ([]byte, error) {
	return json.Marshal(c)
}

func (c *CommentStat) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, c)
}
