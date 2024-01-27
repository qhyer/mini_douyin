package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	ID         int64          `json:"id" gorm:"column:id;primaryKey"`
	FromUserId int64          `json:"from_user_id" gorm:"column:from_user_id;index:user_id_idx,priority:1"`
	ToUserId   int64          `json:"to_user_id" gorm:"column:to_user_id;index:user_id_idx,priority:2"`
	Content    string         `json:"content" gorm:"column:content"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
