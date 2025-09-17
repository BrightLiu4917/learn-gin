package Route

import "github.com/gin-gonic/gin"

var routeInstance *gin.Engine

func InitRoute() *gin.Engine {
	routeInstance = gin.Default()
	return routeInstance
}
