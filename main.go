package main

import (
	"github.com/mashmorsik/rand-quote-grpc/internal/server"
	"github.com/mashmorsik/rand-quote-grpc/logger"
)

func main() {
	logger.Logger()

	server.StartServer()
}
