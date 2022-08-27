package main

import (
	"log"
	"math"
	"net"
	"user-service/config"
	"user-service/proto/auth"

	"google.golang.org/grpc"
)

var conf config.Config

func init() {
	conf = config.LoadConfig()
}

func gRPCListen() {
	log.Printf("Starting gRPC Server at %v", conf.WebPort)
	listener, err := net.Listen("tcp", conf.WebPort)
	if err != nil {
		log.Fatalf("Failed to listen on API Gateway Service: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.MaxConcurrentStreams(math.MaxUint32))

	auth.RegisterAuthServiceServer(grpcServer, &AuthServer{Models: conf.Models})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server on API Gateway Service: %v", err)
	}
}

func main() {
	log.Println("Starting User Service")
	gRPCListen()
}
