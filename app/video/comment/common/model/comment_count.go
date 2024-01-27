package model

import "gorm.io/gorm"

type CommentCount struct {
	gorm.Model
	VideoId      int64 `json:"video_id" gorm:"column:video_id;primaryKey'"`
	CommentCount int64 `json:"comment_count" gorm:"column:comment_count"`
}
