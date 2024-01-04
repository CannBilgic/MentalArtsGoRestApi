package routes

import (
	controller "mentalartsgo/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", controller.ListUsers)
	v1.POST("/user", controller.CreateUser)
	v1.DELETE("/user/:id", controller.DeleteUser)

}
