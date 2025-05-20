package controller

import (
	"go-crm-api-shopdev/internal/service"
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{
		"message": "getUser",
		"info":    uc.userService.GetInfoUser(),
	})
}
