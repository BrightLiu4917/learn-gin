package Services

import (
	"errors"
	"fmt"
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

func (u *UserService) UpdateUser(id int, user *Models.User) error {
	// 1. 验证用户ID有效性
	if id <= 0 {
		return fmt.Errorf("无效的用户ID: %d", id)
	}

	// 2. 检查用户是否存在
	var existingUser Models.User
	if err := u.db.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在 (ID: %d)", id)
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	// 3. 防止ID被意外更新
	user.Id = 0 // 防止更新ID字段

	// 4. 执行更新操作
	result := u.db.Model(&Models.User{}).Where("id = ?", id).Updates(user)

	// 5. 检查更新结果
	if result.Error != nil {
		return fmt.Errorf("更新用户失败: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("未更新任何记录 (ID: %d)", id)
	}

	return nil
}

func (u *UserService) DeleteUser(id int) error {
	result := u.db.Where("id = ?", id).Delete(&Models.User{})
	return result.Error
}
