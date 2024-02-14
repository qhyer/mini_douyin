package mapper

import (
	"douyin/app/user/relation/common/event"
	po "douyin/app/user/relation/common/model"
	"fmt"
)

func RelationActionToPO(rel *event.RelationAction) (*po.Relation, error) {
	if rel == nil {
		return &po.Relation{}, fmt.Errorf("relation is nil")
	}
	return &po.Relation{
		ID:         rel.ID,
		FromUserId: rel.FromUserId,
		ToUserId:   rel.ToUserId,
	}, nil
}
