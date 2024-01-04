package routes

import (
	controller "mentalartsgo/handlers"

	"github.com/gin-gonic/gin"
)

func MonsterRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/monster", controller.ListMonster)
	v1.POST("/monster", controller.CreateMonster)
	v1.DELETE("/monster/:id", controller.DeleteMonster)

}
