package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	UserId    int64          `json:"user_id"  gorm:"column:user_id"`
	VideoId   int64          `json:"video_id" gorm:"column:video_id;index:video_id_created_at_idx,priority:1"`
	Content   string         `json:"content" gorm:"column:content"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;index:video_id_created_at_idx,priority:2"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
