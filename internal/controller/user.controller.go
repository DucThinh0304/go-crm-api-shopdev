package controller

import (
	"go-crm-api-shopdev/internal/service"
	"go-crm-api-shopdev/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetInfoUser(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"abc", "123", "abc"})
}
