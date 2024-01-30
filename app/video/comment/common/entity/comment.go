package entity

import (
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	VideoId   int64     `json:"video_id"`
	User      *User     `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
