package entity

type User struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
	Avatar            string `json:"avatar"`
	BackgroundImage   string `json:"background_image"`
	Signature         string `json:"signature"`
}
