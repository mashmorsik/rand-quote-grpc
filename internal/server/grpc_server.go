package server

import (
	"fmt"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/rand-quote-grpc"
	"github.com/mashmorsik/rand-quote-grpc/internal/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8088))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	qs := handlers.QuotesServer{}
	reflection.Register(grpcServer)

	randquotev1.RegisterRandQuotesServer(grpcServer, qs)
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(fmt.Errorf("%s\n", err))
	}

	var dialOpts []grpc.DialOption
	conn, err := grpc.Dial("localhost:8088", dialOpts...)
	if err != nil {
		fmt.Println(fmt.Errorf("%s\n", err))
	}
	defer conn.Close()
}
