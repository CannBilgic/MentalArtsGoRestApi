package routes

import (
	controller "mentalartsgo/handlers"

	"github.com/gin-gonic/gin"
)

func BeybladeRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	v1.GET("/beyblade", controller.ListBeyblade)
	v1.POST("/beyblade", controller.CreateBeyblade)
	v1.DELETE("beyblade/:id", controller.DeleteBeyblade)

}
