// Package Routes 注册路由
package Routes

import (
	"learn-gin/Api/Controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/api/v1")
	{
		userController := Controllers.NewUserController()
		// 注册一个路由
		v1.POST("/create-user", userController.CreateUser)
		v1.GET("/list-user", userController.GetAllUser)
		v1.GET("/first-user", userController.FirstUser)
	}
}
