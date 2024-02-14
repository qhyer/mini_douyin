package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                int64          `json:"id" gorm:"primaryKey;column:id"`
	Name              string         `json:"name" gorm:"column:name;index:name_idx"`
	EncryptedPassword string         `json:"encrypted_password" gorm:"column:password"`
	CreatedAt         time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
