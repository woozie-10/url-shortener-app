package main

import (
	"fmt"
	"log"
	"url-shortener-app/config"
	_ "url-shortener-app/docs"
	"url-shortener-app/router"
)

// @title URL Shortener App
// @version 1.0
// @description This web application serves as a URL Shortener, built with Go and Gin. Users can easily shorten URLs through a user-friendly interface.
// @host localhost:2311
// @BasePath /
// @schemes http
func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("Config initialization error: %s", err.Error())
	}
	port := config.Config.GetString("server.port")
	router.InitRouter()
	if err := router.Router.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		log.Fatalf("Server startup error: %s", err.Error())
	}
}
