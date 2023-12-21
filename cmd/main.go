package main

import (
	"fmt"
	"log"
	"url-shortener-app/config"
	"url-shortener-app/router"
)

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
