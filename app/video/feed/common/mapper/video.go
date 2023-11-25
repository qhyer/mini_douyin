package mapper

import (
	dto "douyin/api/video/publish/service/v1"
	do "douyin/app/video/feed/common/entity"
	"time"
)

func VideoFromDTO(d *dto.VideoInfo) (*do.Video, error) {
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

func VideoFromDTOs(dtos []*dto.VideoInfo) ([]*do.Video, error) {
	videos := make([]*do.Video, 0, len(dtos))
	for _, d := range dtos {
		v, err := VideoFromDTO(d)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	return videos, nil
}
