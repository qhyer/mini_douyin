package entity

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	VideoId   int64     `json:"video_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
