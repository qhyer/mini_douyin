package entity

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID         int64     `json:"id"`
	FromUserId int64     `json:"from_user_id"`
	ToUserId   int64     `json:"to_user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func (m *Message) MarshalJson() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) UnmarshalJson(b []byte) error {
	return json.Unmarshal(b, m)
}
