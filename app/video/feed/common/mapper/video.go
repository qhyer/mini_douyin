package mapper

import (
	"time"

	feedDTO "douyin/api/video/feed/service/v1"
	publishDTO "douyin/api/video/publish/service/v1"
	do "douyin/app/video/feed/common/entity"
)

func VideoFromPublishDTO(d *publishDTO.VideoInfo) (*do.Video, error) {
	if d == nil {
		return &do.Video{}, nil
	}
	return &do.Video{
		ID: d.GetId(),
		User: &do.User{
			ID: d.GetAuthorId(),
		},
		PlayURL:   d.GetPlayUrl(),
		CoverURL:  d.GetCoverUrl(),
		Title:     d.GetTitle(),
		CreatedAt: time.Unix(d.GetCreateTime(), 0),
	}, nil
}

func VideoFromPublishDTOs(dtos []*publishDTO.VideoInfo) ([]*do.Video, error) {
	if dtos == nil {
		return []*do.Video{}, nil
	}
	videos := make([]*do.Video, 0, len(dtos))
	for _, d := range dtos {
		v, err := VideoFromPublishDTO(d)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}

func VideoToFeedDTO(v *do.Video) (*feedDTO.Video, error) {
	if v == nil {
		return &feedDTO.Video{}, nil
	}
	return &feedDTO.Video{
		Id: v.ID,
		Author: &feedDTO.User{
			Id:       v.User.ID,
			Name:     v.User.Name,
			IsFollow: v.User.IsFollow,
			Avatar:   &v.User.Avatar,
		},
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}, nil
}

func VideoToFeedDTOs(vs []*do.Video) ([]*feedDTO.Video, error) {
	if vs == nil {
		return []*feedDTO.Video{}, nil
	}
	videos := make([]*feedDTO.Video, 0, len(vs))
	for _, v := range vs {
		v, err := VideoToFeedDTO(v)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}
