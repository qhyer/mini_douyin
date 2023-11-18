package entity

import (
	"encoding/json"
	"time"
)

type Video struct {
	ID            int64     `json:"id"`
	AuthorID      int64     `json:"author_id"`
	Title         string    `json:"title"`
	VideoFileName string    `json:"video_file_name"`
	PlayURL       string    `json:"play_url"`
	CoverURL      string    `json:"cover_url"`
	CreatedAt     time.Time `json:"created_at"`
}

func (v *Video) MarshalJson() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Video) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, v)
}
