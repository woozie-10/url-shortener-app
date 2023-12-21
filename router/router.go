package router

import (
	"github.com/gin-gonic/gin"
	"url-shortener-app/handlers"
)

var Router *gin.Engine = gin.Default()

func InitRouter() {
	Router.LoadHTMLGlob("frontend/html/*")
	RouterRegister()
}

func RouterRegister() {
	Router.GET("/", handlers.IndexHandler)
	Router.POST("/", handlers.ShortenURLHandler)
}
