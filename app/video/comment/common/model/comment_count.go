package model

import "gorm.io/gorm"

type CommentCount struct {
	gorm.Model
	VideoId      int64 `json:"video_id"`
	CommentCount int64 `json:"comment_count"`
}
