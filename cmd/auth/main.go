package main

import (
	server2 "github.com/mashmorsik/rand-quote-grpc/infrastructure/auth/server"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"
)

func main() {
	log.Logger()

	server2.StartAuthServer()
}
