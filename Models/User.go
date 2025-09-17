package Models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	RealName  string    `json:"real_name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"-" gorm:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
}
