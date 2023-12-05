package model

import (
	"gorm.io/gorm"
	"time"
)

type Relation struct {
	gorm.Model
	ID         int64          `json:"id"`
	FromUserId int64          `json:"from_user_id"`
	ToUserId   int64          `json:"to_user_id"`
	Attr       uint32         `json:"attr"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
