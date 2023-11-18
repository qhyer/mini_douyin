package mapper

import (
	do "douyin/app/video/comment/common/entity"
	po "douyin/app/video/comment/common/model"
)

func CommentFromPO(c *po.Comment) (*do.Comment, error) {
	return &do.Comment{
		ID:        c.ID,
		UserId:    c.UserId,
		VideoId:   c.VideoId,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}, nil
}

func CommentFromPOs(c []*po.Comment) ([]*do.Comment, error) {
	res := make([]*do.Comment, 0, len(c))
	for _, v := range c {
		res = append(res, &do.Comment{
			ID:        v.ID,
			UserId:    v.UserId,
			VideoId:   v.VideoId,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func CommentToPOs(c []*do.Comment) ([]*po.Comment, error) {
	res := make([]*po.Comment, 0, len(c))
	for _, v := range c {
		res = append(res, &po.Comment{
			ID:        v.ID,
			UserId:    v.UserId,
			VideoId:   v.VideoId,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}
