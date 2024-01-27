package model

import "gorm.io/gorm"

type UserFavoriteCount struct {
	gorm.Model
	UserId         int64 `gorm:"column:user_id;primaryKey" json:"user_id"`
	FavoriteCount  int64 `gorm:"column:favorite_count" json:"favorite_count"`
	FavoritedCount int64 `gorm:"column:favorited_count" json:"favorited_count"`
}
