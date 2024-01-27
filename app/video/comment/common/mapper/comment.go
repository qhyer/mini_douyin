package mapper

import (
	v1 "douyin/api/video/comment/service/v1"
	do "douyin/app/video/comment/common/entity"
	po "douyin/app/video/comment/common/model"
	"fmt"
)

func CommentFromPO(c *po.Comment) (*do.CommentAction, error) {
	return &do.CommentAction{
		ID: c.ID,
		User: &do.User{
			ID: c.UserId,
		},
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}, nil
}

func CommentFromPOs(c []*po.Comment) ([]*do.CommentAction, error) {
	res := make([]*do.CommentAction, 0, len(c))
	for _, v := range c {
		res = append(res, &do.CommentAction{
			ID: v.ID,
			User: &do.User{
				ID: v.UserId,
			},
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func CommentToPO(c *do.CommentAction) (*po.Comment, error) {
	return &po.Comment{
		ID:        c.ID,
		UserId:    c.User.ID,
		VideoId:   c.VideoId,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}, nil
}

func CommentToPOs(c []*do.CommentAction) ([]*po.Comment, error) {
	res := make([]*po.Comment, 0, len(c))
	for _, v := range c {
		res = append(res, &po.Comment{
			ID:        v.ID,
			UserId:    v.User.ID,
			VideoId:   v.VideoId,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func CommentToDTO(comment *do.CommentAction) (*v1.CommentInfo, error) {
	return &v1.CommentInfo{
		Id: comment.ID,
		User: &v1.User{
			Id:     comment.User.ID,
			Name:   comment.User.Name,
			Avatar: &comment.User.Avatar,
		},
		Content:    comment.Content,
		CreateDate: fmt.Sprintf("%d-%d", comment.CreatedAt.Month(), comment.CreatedAt.Day()),
	}, nil
}

func CommentToDTOs(comments []*do.CommentAction) ([]*v1.CommentInfo, error) {
	res := make([]*v1.CommentInfo, 0, len(comments))
	for _, v := range comments {
		res = append(res, &v1.CommentInfo{
			Id:         v.ID,
			Content:    v.Content,
			CreateDate: fmt.Sprintf("%d-%d", v.CreatedAt.Month(), v.CreatedAt.Day()),
		})
	}
	return res, nil
}
