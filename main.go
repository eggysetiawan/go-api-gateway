package main

import (
	"github.com/eggysetiawan/go-api-gateway/app"
	"github.com/eggysetiawan/go-api-gateway/logger"
)

func main() {
	logger.Info("Starting Application...")
	app.Start()

}
