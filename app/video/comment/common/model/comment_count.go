package model

type CommentCount struct {
	VideoId      int64 `json:"video_id" gorm:"column:video_id;primaryKey'"`
	CommentCount int64 `json:"comment_count" gorm:"column:comment_count"`
}
