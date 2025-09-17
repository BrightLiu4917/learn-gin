package Db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 声明全局变量
var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

// Connect 初始化数据库连接 (单例模式)
func Connect() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = gorm.Open("mysql", "root:root@(10.123.234.48:3306)/learn-gin?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			panic("数据库连接失败: " + err.Error())
		}

		// 配置连接池
		dbInstance.DB().SetMaxIdleConns(10)
		dbInstance.DB().SetMaxOpenConns(100)
	})
	return dbInstance
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	if dbInstance == nil {
		panic("数据库未初始化，请先调用 Db.Connect()")
	}
	return dbInstance
}

// Close 关闭数据库连接
func Close() {
	if dbInstance != nil {
		_ = dbInstance.Close()
	}
}
