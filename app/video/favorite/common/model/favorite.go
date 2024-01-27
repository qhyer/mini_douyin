package model

import (
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	gorm.Model
	ID        int64          `gorm:"column:id;primaryKey" json:"id"`
	UserId    int64          `gorm:"column:user_id;idx:user_id_idx" json:"user_id"`
	VideoId   int64          `gorm:"column:video_id" json:"video_id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
