package routes

import (
	"go-auth/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/signup", controller.Signup)
	r.GET("/home", controller.Home)
	r.GET("/premium", controller.Premium)
	r.GET("/logout", controller.Logout)
}
