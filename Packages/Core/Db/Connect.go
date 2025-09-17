package Db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func Connect() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = gorm.Open("mysql", "root:root@(local-mysql8:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}

		// 可选：配置连接池参数
		dbInstance.DB().SetMaxIdleConns(10)
		dbInstance.DB().SetMaxOpenConns(100)
	})
	return dbInstance
}
