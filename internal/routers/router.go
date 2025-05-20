package routers

import (
	c "go-crm-api-shopdev/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1/2025")
	v1.GET("/ping", c.NewPongController().Pong)
	v1.GET("/user", c.NewUserController().GetInfoUser)
	return r
}
