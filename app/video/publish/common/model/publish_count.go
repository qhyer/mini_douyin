package model

type PublishCount struct {
	UserID     int64 `json:"user_id" gorm:"primaryKey;column:user_id"`
	VideoCount int   `json:"video_count" gorm:"column:video_count"`
}
