package model

import "gorm.io/gorm"

type UserFavoriteCount struct {
	gorm.Model
	UserId         int64
	FavoriteCount  int64
	FavoritedCount int64
}
