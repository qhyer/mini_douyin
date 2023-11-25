package entity

import "encoding/json"

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	Message         string `json:"message"`
	MsgType         int64  `json:"msg_type"`
}

func (u *User) MarshalJson() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, u)
}
