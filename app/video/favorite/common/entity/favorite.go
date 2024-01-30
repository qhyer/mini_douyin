package event

import (
	"encoding/json"
	"time"
)

type Favorite struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	VideoId   int64     `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (f *Favorite) MarshalJson() ([]byte, error) {
	return json.Marshal(f)
}

func (f *Favorite) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, f)
}
