package model

import (
	"time"

	"gorm.io/gorm"
)

type Relation struct {
	ID         int64          `json:"id" gorm:"primaryKey"`
	FromUserId int64          `json:"from_user_id" gorm:"column:from_user_id;index:user_id_idx,priority:1"`
	ToUserId   int64          `json:"to_user_id" gorm:"column:to_user_id;index:user_id_idx,priority:2"`
	Attr       uint32         `json:"attr" gorm:"column:attr"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
