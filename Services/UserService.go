package Services

import (
	"github.com/jinzhu/gorm"
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
