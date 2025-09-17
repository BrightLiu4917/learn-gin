package Services

import (
	"github.com/jinzhu/gorm"
	"learn-gin/Models"
	"learn-gin/Packages/Core/Db"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: Db.GetDB(),
	}
}

func (u *UserService) CreateUser(user *Models.User) error {
	return u.db.Create(&user).Error
}
