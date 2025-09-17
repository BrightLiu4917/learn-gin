package Controllers

import (
	"learn-gin/Models"
	"learn-gin/Packages/Tools/Response"
	"learn-gin/Services"

	"github.com/gin-gonic/gin"
)

var user Models.User

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

	// 解析请求体中的JSON数据
	if err := ctx.ShouldBindJSON(&user); err != nil {
		Response.InvalidParams(ctx, "")
		return
	}

	// 调用服务创建用户
	if err := c.service.CreateUser(&user); err != nil {
		Response.Error(ctx, 500, "用户创建失败")
		return
	}

	Response.Success(ctx, user)
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	// 1. 调用服务层获取所有用户
	users, err := c.service.GetAllUser()
	if err != nil {
		// 2. 错误处理：返回500状态码和错误信息
		Response.Error(ctx, 500, "获取失败")
		return
	}

	// 3. 成功响应：返回200状态码和用户数据
	Response.Success(ctx, users)
}

func (c *UserController) FirstUser(ctx *gin.Context) {
	users, err := c.service.FirstUser(ctx.Query("id"))
	if err != nil {
		Response.Error(ctx, 500, "获取失败")
		return
	}
	Response.Success(ctx, users)
}
