package server

import (
	"fmt"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/auth-grpc"
	data2 "github.com/mashmorsik/rand-quote-grpc/infrastructure/auth/data"
	auth2 "github.com/mashmorsik/rand-quote-grpc/internal/auth"
	"github.com/mashmorsik/rand-quote-grpc/internal/handlers/auth"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

func StartAuthServer() {
	data := data2.NewData(data2.MustConnect())
	log.Infof("database is set %#+v", data.Db())

	var port = fmt.Sprintf("localhost:%d", 8090)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	authServer := auth.NewAuthServer(auth2.NewUserData(*data))

	reflection.Register(grpcServer)

	randquotev1.RegisterAuthServer(grpcServer, authServer)
	log.Infof("grpc_server started on: %s", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Nlog.Err(err).Msg("failed to start auth server")
		return
	}
}
