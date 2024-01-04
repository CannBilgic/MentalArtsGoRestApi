package main

import (
	"mentalartsgo/config"
	"mentalartsgo/docs"
	"mentalartsgo/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1/"
	config.Connect()
	routes.UserRoute(router)
	routes.MonsterRoute(router)
	routes.BeybladeRoute(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":3000")
}
