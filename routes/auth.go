package routes

import (
	"github.com/Emrul-Hasan-Emon/jwt-go/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.SignUp())
	router.POST("users/login", controllers.LogIn())
}
