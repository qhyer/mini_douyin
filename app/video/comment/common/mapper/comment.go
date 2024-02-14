package mapper

import (
	"fmt"

	"douyin/app/video/comment/common/event"

	v1 "douyin/api/video/comment/service/v1"
	do "douyin/app/video/comment/common/entity"
	po "douyin/app/video/comment/common/model"
)

func CommentFromPO(c *po.Comment) (*do.Comment, error) {
	if c == nil {
		return &do.Comment{}, fmt.Errorf("comment is nil")
	}
	return &do.Comment{
		ID: c.ID,
		User: &do.User{
			ID: c.UserId,
		},
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}, nil
}

func CommentFromPOs(c []*po.Comment) ([]*do.Comment, error) {
	res := make([]*do.Comment, 0, len(c))
	for _, v := range c {
		res = append(res, &do.Comment{
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

func CommentToPO(c *do.Comment) (*po.Comment, error) {
	if c == nil {
		return &po.Comment{}, fmt.Errorf("comment is nil")
	}
	return &po.Comment{
		ID:        c.ID,
		UserId:    c.User.ID,
		VideoId:   c.VideoId,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}, nil
}

func CommentToPOs(c []*do.Comment) ([]*po.Comment, error) {
	res := make([]*po.Comment, 0, len(c))
	if c == nil {
		return res, fmt.Errorf("comment is nil")
	}
	for _, v := range c {
		if v == nil {
			continue
		}
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

func CommentToDTO(comment *do.Comment) (*v1.CommentInfo, error) {
	if comment == nil {
		return &v1.CommentInfo{}, fmt.Errorf("comment is nil")
	}
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

func CommentToDTOs(comments []*do.Comment) ([]*v1.CommentInfo, error) {
	res := make([]*v1.CommentInfo, 0, len(comments))
	if comments == nil {
		return res, fmt.Errorf("comments is nil")
	}
	for _, v := range comments {
		if v == nil {
			continue
		}
		res = append(res, &v1.CommentInfo{
			Id:         v.ID,
			Content:    v.Content,
			CreateDate: fmt.Sprintf("%d-%d", v.CreatedAt.Month(), v.CreatedAt.Day()),
		})
	}
	return res, nil
}

func ParseCommentFromCommentAction(commentAction *event.CommentAction) (*do.Comment, error) {
	if commentAction == nil {
		return &do.Comment{}, fmt.Errorf("commentAction is nil")
	}
	return &do.Comment{
		ID:      commentAction.ID,
		VideoId: commentAction.VideoId,
		User: &do.User{
			ID: commentAction.UserId,
		},
		Content:   commentAction.Content,
		CreatedAt: commentAction.CreatedAt,
	}, nil
}

func ParseCommentFromCommentActions(commentActions []*event.CommentAction) ([]*do.Comment, error) {
	res := make([]*do.Comment, 0, len(commentActions))
	if commentActions == nil {
		return []*do.Comment{}, fmt.Errorf("commentActions is nil")
	}
	for _, v := range commentActions {
		if v == nil {
			continue
		}
		res = append(res, &do.Comment{
			ID:      v.ID,
			VideoId: v.VideoId,
			User: &do.User{
				ID: v.UserId,
			},
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}
