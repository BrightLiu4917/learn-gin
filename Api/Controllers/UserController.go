package Controllers

import (
	"learn-gin/Models"
	"learn-gin/Packages/Tools/Response"
	"learn-gin/Services"
	"log"
	"strings"

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

func (c *UserController) UpdateUser(ctx *gin.Context) {
	// 1. 获取并验证ID

	// 2. 绑定请求数据
	var user Models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("请求数据绑定失败: %v", err)
		Response.Error(ctx, 400, "无效的请求数据")
		return
	}
	id := user.Id

	// 3. 记录请求日志
	log.Printf("更新用户请求 | ID: %d | 数据: %+v", id, user)

	// 4. 调用服务层
	if err := c.service.UpdateUser(user.Id, &user); err != nil {
		// 根据错误类型返回不同状态码
		if strings.Contains(err.Error(), "用户不存在") {
			Response.Error(ctx, 404, err.Error())
		} else if strings.Contains(err.Error(), "无效的用户ID") {
			Response.Error(ctx, 400, err.Error())
		} else {
			log.Printf("更新用户错误: %v", err)
			Response.Error(ctx, 500, "更新用户失败")
		}
		return
	}

	// 5. 成功响应
	Response.Success(ctx, gin.H{})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {

	var user Models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("请求数据绑定失败: %v", err)
		Response.Error(ctx, 400, "无效的请求数据")
		return
	}
	id := user.Id
	u.service.DeleteUser(id)
}
