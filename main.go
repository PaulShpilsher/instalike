package main

import (
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/webserver"
)

// @title InstaLike API
// @version 2.0
// @description This is a instagram-like server (instalike).
// @termsOfService http://swagger.io/terms/

// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	config := config.LoadConfig()
	server := webserver.NewWebServer(&config)
	server.Start()
}
