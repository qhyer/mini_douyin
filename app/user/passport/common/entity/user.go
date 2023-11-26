package entity

import "encoding/json"

type User struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
	Avatar            string `json:"avatar"`
	BackgroundImage   string `json:"background_image"`
	Signature         string `json:"signature"`
}

func (u *User) MarshalJson() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalJson(b []byte) error {
	return json.Unmarshal(b, u)
}
