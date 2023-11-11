package entity

import "time"

type Video struct {
	ID            int64     `json:"id"`
	User          User      `json:"user"`
	PlayURL       string    `json:"play_url"`
	CoverURL      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"created_at"`
}
