package main

import (
	"github.com/pasDamola/go-microservices/app"
	"github.com/pasDamola/go-microservices/logger"
)


func main() {

	logger.Info("Starting the application...")
	app.Start()
	
}

