package server

import (
	"fmt"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/rand-quote-grpc"
	"github.com/mashmorsik/rand-quote-grpc/internal/handlers/quote"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func StartServer() {
	var port = fmt.Sprintf("localhost:%d", 8088)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	randquotev1.RegisterRandQuotesServer(grpcServer, quote.QuotesServer{})
	log.Nlog.Info().Msgf("grpc_server started on: %s", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Nlog.Err(err).Msg("failed to start server")
		return
	}

}
