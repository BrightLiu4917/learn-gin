package Controllers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/Models"
	"learn-gin/Services"
)

type UserController struct {
	service Services.UserService
}

func NewUserController() *UserController {
	userService := Services.NewUserService()
	if userService == nil {
		// 处理服务初始化失败的情况
		panic("failed to initialize user service")
	}
	return &UserController{
		service: *userService,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user Models.User

	// 解析请求体中的JSON数据
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request format: " + err.Error(),
		})
		return
	}

	// 调用服务创建用户
	if err := c.service.CreateUser(&user); err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User created successfully",
	})
}
