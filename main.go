package main

import (
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/webserver"
)

func main() {

	config := config.LoadConfig()
	server := webserver.NewWebServer(&config)
	server.Start()
}
