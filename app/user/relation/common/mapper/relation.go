package mapper

import (
	do "douyin/app/user/relation/common/entity"
	po "douyin/app/user/relation/common/model"
)

func RelationActionToPO(rel *do.RelationAction) (*po.Relation, error) {
	return &po.Relation{
		ID:         rel.ID,
		FromUserId: rel.FromUserId,
		ToUserId:   rel.ToUserId,
	}, nil
}
