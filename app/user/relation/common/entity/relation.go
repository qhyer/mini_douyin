package entity

import (
	"encoding/json"
	"time"
)

type RelationType uint32

const (
	NoRelation RelationType = iota + 1
	RelationFollowed
	RelationFollowing
	RelationFriend
)

const (
	RelationAttrFollowing = 1 << iota
	RelationAttrFollowed
)

type Relation struct {
	ID         int64        `json:"id"`
	Type       RelationType `json:"type"`
	FromUserId int64        `json:"from_user_id"`
	ToUserId   int64        `json:"to_user_id"`
	CreatedAt  time.Time    `json:"created_at"`
}

func (r *Relation) MarshalJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Relation) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, r)
}

func Attr(attribute uint32) RelationType {
	if attribute&RelationAttrFollowed > 0 && attribute&RelationAttrFollowing > 0 {
		return RelationFriend
	}
	if attribute&RelationAttrFollowing > 0 {
		return RelationAttrFollowing
	}
	if attribute&RelationAttrFollowed > 0 {
		return RelationFollowed
	}
	return NoRelation
}

func SetAttr(attribute uint32, mask uint32) uint32 {
	return attribute | mask
}

func UnsetAttr(attribute uint32, mask uint32) uint32 {
	return attribute & ^mask
}
