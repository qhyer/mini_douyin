package event

import (
	"encoding/json"
	"time"
)

type VideoUpload struct {
	ID            int64     `json:"id"`
	AuthorID      int64     `json:"author_id"`
	Title         string    `json:"title"`
	VideoFileName string    `json:"video_file_name"`
	CreatedAt     time.Time `json:"created_at"`
}

func (v *VideoUpload) MarshalJson() ([]byte, error) {
	return json.Marshal(v)
}

func (v *VideoUpload) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, v)
}
