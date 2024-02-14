package model

type VideoFavoritedCount struct {
	VideoId        int64 `gorm:"column:video_id;primaryKey" json:"video_id"`
	FavoritedCount int64 `gorm:"column:favorited_count" json:"favorited_count"`
}
