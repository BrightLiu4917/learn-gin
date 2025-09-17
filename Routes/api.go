// Package Routes 注册路由
package Routes

import (
	"github.com/gin-gonic/gin"
	"learn-gin/Api/Controllers"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		userController := Controllers.NewUserController()
		// 注册一个路由
		v1.GET("/create-user", userController.CreateUser)
	}
}
