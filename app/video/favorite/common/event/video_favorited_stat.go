package event

import "encoding/json"

type VideoFavoritedStat struct {
	VideoId int64 `json:"video_id"`
	Delta   int64 `json:"delta"`
}

func (v *VideoFavoritedStat) MarshalJson() ([]byte, error) {
	return json.Marshal(v)
}

func (v *VideoFavoritedStat) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, v)
}
