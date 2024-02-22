package main

import (
	"github.com/mashmorsik/rand-quote-grpc/infrastructure/quote/server"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"
)

func main() {
	log.Logger()

	server.StartServer()
}
