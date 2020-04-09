package main

import (
	"fmt"
	"github.com/claudioed/grpc-service/internal/server"
	"net"
)

func main() {
	cfg := server.NewConfig(server.NewLog())
	srv, _ := server.NewGRPCServer(cfg)
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
	srv.Serve(lis)
}
