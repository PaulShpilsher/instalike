package main

import (
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/webserver"
)

// @title InstaLike API
// @version 2.0
// @description This is a instagram-like server (instalike).
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {

	config := config.LoadConfig()
	server := webserver.NewWebServer(&config)
	server.Start()
}
