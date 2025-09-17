package Init

import (
	"learn-gin/Packages/Core/Db"
	"learn-gin/Packages/Core/Route"
	"learn-gin/Routes"
)

func Bootstrap() {
	// 1. 初始化数据库等基础组件
	Db.GetDB()

	// 2. 获取路由实例并注册路由
	router := Route.GetRouter()
	Routes.RegisterAPIRoutes(router)
}
