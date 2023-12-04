package model

import "gorm.io/gorm"

type VideoFavoritedCount struct {
	gorm.Model
	VideoId        int64
	FavoritedCount int64
}
