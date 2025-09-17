package Services

import (
	"github.com/jinzhu/gorm"
	"learn-gin/Models"
	"learn-gin/Packages/Core/Db"
)

// 用户服务
type UserService struct {
	db *gorm.DB
}

// 构造函数
func NewUserService() *UserService {
	return &UserService{
		db: Db.GetDB(),
	}
}

// 创建用户的方法
func (u *UserService) CreateUser(user *Models.User) error {
	return u.db.Create(&user).Error
}
