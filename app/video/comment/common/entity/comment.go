package entity

import (
	"encoding/json"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	VideoId   int64     `json:"video_id"`
	User      *User     `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentActionType int

const (
	CommentActionPublish CommentActionType = iota + 1
	CommentActionDelete
)

type CommentAction struct {
	ID        int64             `json:"id"`
	Type      CommentActionType `json:"type"`
	UserId    int64             `json:"user_id"`
	VideoId   int64             `json:"video_id"`
	Content   string            `json:"content"`
	CreatedAt time.Time         `json:"created_at"`
}

func (c *CommentAction) MarshalJson() ([]byte, error) {
	return json.Marshal(c)
}

func (c *CommentAction) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, c)
}
