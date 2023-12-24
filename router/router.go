package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"url-shortener-app/handlers"

	swaggerFiles "github.com/swaggo/files"
)

var Router *gin.Engine = gin.Default()

func InitRouter() {
	Router.LoadHTMLGlob("frontend/html/*")
	RouterRegister()
}

func RouterRegister() {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/", handlers.IndexHandler)
	Router.POST("/", handlers.ShortenURLHandler)
}
