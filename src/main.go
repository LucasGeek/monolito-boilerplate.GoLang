package main

import (
	"server/src/commons/config"
	"server/src/layers/app/api"
	"server/src/layers/app/di"
)

func main() {
	cfg := config.LoadConfig()

	container := di.InitializeContainer()
	server := api.NewFiberServer(container)
	server.SetupRoutes()
	server.Run(cfg.Port)
}
