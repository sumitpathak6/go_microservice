package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/sumitpathak6/go_microservice/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("Failed to start app:", err)
	}

}

//!DOCKER RUNNING COMMMAND  docker run -p 6379:6379 redis:latest
