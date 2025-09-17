package Services

import (
	"learn-gin/Models"
	"learn-gin/Packages/Core/Db"

	"github.com/jinzhu/gorm"
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

func (u *UserService) GetAllUser() ([]Models.User, error) {
	var users []Models.User
	result := u.db.Find(&users)
	return users, result.Error
}

func (u *UserService) FirstUser(id string) (Models.User, error) {
	var user Models.User
	result := u.db.First(&user, id)
	return user, result.Error
}
