package routes

import (
	"github.com/Emrul-Hasan-Emon/jwt-go/controllers"
	"github.com/Emrul-Hasan-Emon/jwt-go/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middlewares.Authenticate())
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/users/:user_id", controllers.GetUserById())
}
