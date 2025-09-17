package Models

import (
	"time"
)

// 用户模型
type User struct {
	Id        uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	RealName  string    `gorm:"column:real_name;type:varchar(32);NOT NULL" json:"real_name"`
	UserName  string    `gorm:"column:user_name;type:varchar(32);NOT NULL" json:"user_name"`
	Password  string    `gorm:"column:password;type:varchar(64);NOT NULL" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;default:null" json:"deleted_at"`
}
