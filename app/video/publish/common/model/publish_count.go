package model

import "gorm.io/gorm"

type PublishCount struct {
	gorm.Model
	UserID int64 `json:"user_id"`
	Count  int64 `json:"count"`
}
