package entity

import (
	"encoding/json"
	"time"
)

type RelationActionType int

const (
	RelationActionFollow RelationActionType = iota + 1
	RelationActionUnFollow
)

type Relation struct {
	ID         int64              `json:"id"`
	Type       RelationActionType `json:"type"`
	FromUserId int64              `json:"from_user_id"`
	ToUserId   int64              `json:"to_user_id"`
	CreatedAt  time.Time          `json:"created_at"`
}

func (r *Relation) MarshalJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Relation) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, r)
}
