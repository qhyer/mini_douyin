package entity

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsFollow bool   `json:"is_follow"`
	Avatar   string `json:"avatar"`
}
