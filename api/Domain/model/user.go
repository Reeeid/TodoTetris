package model

import "time"

type (
	User struct {
		Username     string `gorm:"primaryKey" json:"username"` // IDの代わりに名前を主キーに
		PasswordHash string `json:"-"`                          // パスワードは絶対返さない
		CreatedAt    time.Time
	}
	RegisterUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	LoginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
