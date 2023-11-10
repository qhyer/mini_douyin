package mapper

import (
	dto "douyin/api/video/publish/service/v1"
	do "douyin/app/video/publish/common/entity"
	po "douyin/app/video/publish/common/model"
)

func VideoFromPO(v *po.Video) (*do.Video, error) {
	return &do.Video{
		ID:        v.ID,
		AuthorID:  v.AuthorID,
		Title:     v.Title,
		PlayURL:   v.PlayURL,
		CoverURL:  v.CoverURL,
		CreatedAt: v.CreatedAt,
	}, nil
}

func VideoToPO(v *do.Video) (*po.Video, error) {
	return &po.Video{
		ID:        v.ID,
		AuthorID:  v.AuthorID,
		Title:     v.Title,
		PlayURL:   v.PlayURL,
		CoverURL:  v.CoverURL,
		CreatedAt: v.CreatedAt,
	}, nil
}

func VideoFromPOs(v []*po.Video) ([]*do.Video, error) {
	var videos []*do.Video
	for _, item := range v {
		video, err := VideoFromPO(item)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func VideoToDTO(v *do.Video) (*dto.VideoInfo, error) {
	return &dto.VideoInfo{
		Id:         v.ID,
		AuthorId:   v.AuthorID,
		Title:      v.Title,
		CoverUrl:   v.CoverURL,
		PlayUrl:    v.PlayURL,
		CreateTime: v.CreatedAt.UnixMilli(),
	}, nil
}

func VideoToDTOs(vs []*do.Video) ([]*dto.VideoInfo, error) {
	var videos []*dto.VideoInfo
	for _, item := range vs {
		video, err := VideoToDTO(item)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil

}
