package mapper

import (
	do "douyin/app/video/favorite/common/entity"
	po "douyin/app/video/favorite/common/model"
)

func FavoriteFromPO(po *po.Favorite) (*do.Favorite, error) {
	return &do.Favorite{
		ID:        po.ID,
		UserId:    po.UserId,
		VideoId:   po.VideoId,
		CreatedAt: po.CreatedAt,
	}, nil
}

func FavoriteToPO(do *do.Favorite) (*po.Favorite, error) {
	return &po.Favorite{
		ID:        do.ID,
		UserId:    do.UserId,
		VideoId:   do.VideoId,
		CreatedAt: do.CreatedAt,
	}, nil
}

func FavoriteFromPOs(pos []*po.Favorite) ([]*do.Favorite, error) {
	var dos []*do.Favorite
	for _, p := range pos {
		d, err := FavoriteFromPO(p)
		if err != nil {
			return nil, err
		}
		dos = append(dos, d)
	}
	return dos, nil
}

func FavoriteToPOs(dos []*do.Favorite) ([]*po.Favorite, error) {
	var pos []*po.Favorite
	for _, d := range dos {
		p, err := FavoriteToPO(d)
		if err != nil {
			return nil, err
		}
		pos = append(pos, p)
	}
	return pos, nil
}
