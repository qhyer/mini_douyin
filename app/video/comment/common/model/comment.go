package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ID        int64          `json:"id" gorm:"primaryKey"`
	UserId    int64          `json:"user_id"`
	VideoId   int64          `json:"video_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
