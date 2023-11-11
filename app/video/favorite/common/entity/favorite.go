package entity

const (
	FavoriteActionAdd = iota + 1
	FavoriteActionDelete
)

type Favorite struct {
	ID        int64 `json:"id"`
	UserId    int64 `json:"user_id"`
	VideoId   int64 `json:"video_id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
