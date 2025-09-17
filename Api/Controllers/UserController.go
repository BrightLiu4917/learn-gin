package Controllers

import (
	"learn-gin/Models"

	"github.com/gin-gonic/gin"
	"learn-gin/Services"
)

type UserController struct {
	service Services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: *Services.NewUserService(),
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user Models.User

	if err := c.service.CreateUser(&user); err != nil {

	}

}
