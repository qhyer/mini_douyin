package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID        int64          `json:"id" gorm:"primaryKey;column:id"`
	AuthorID  int64          `json:"author_id" gorm:"column:author_id;index:author_id_idx"`
	Title     string         `json:"title" gorm:"column:title"`
	PlayURL   string         `json:"play_url" gorm:"column:play_url"`
	CoverURL  string         `json:"cover_url" gorm:"column:cover_url"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;index:created_at_idx"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (v *Video) MarshalJson() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Video) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, v)
}
