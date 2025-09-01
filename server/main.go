package main

import (
	"log"

	"github.com/yuudi/ero-runner/server/config"
	"github.com/yuudi/ero-runner/server/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(config.GetConfig().Listen); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
