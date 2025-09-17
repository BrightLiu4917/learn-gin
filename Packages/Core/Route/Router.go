package Route

import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	router = gin.Default()
	// 可在此添加全局中间件
}

func GetRouter() *gin.Engine {
	return router
}
