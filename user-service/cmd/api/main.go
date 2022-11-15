package main

import (
	"context"
	"log"
	"math"
	"net"
	"time"
	"user-service/config"
	"user-service/proto/auth"
	"user-service/proto/user"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

var conf config.Config
var connCount uint32

func main() {
	log.Println("Starting User Service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to pSQL DB")
	}

	conf = config.LoadConfig(conn)

	gRPCListen()
}

func openDB(dsn string) (*pgx.Conn, error) {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

// TODO shift to env
func connectToDB() *pgx.Conn {
	dsn := "postgres://admin:password@host.docker.internal:5432/user_db"
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Failed connection attempt to pSQL: %v", err)
			connCount++
		} else {
			log.Println("Successfully connected to pSQL database")
			return connection
		}

		if connCount > 20 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds..")
		time.Sleep(2 * time.Second)
		continue
	}
}

func gRPCListen() {
	log.Printf("Starting gRPC Server at %v", conf.WebPort)
	listener, err := net.Listen("tcp", conf.WebPort)
	if err != nil {
		log.Fatalf("Failed to listen on API Gateway Service: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.MaxConcurrentStreams(math.MaxUint32))

	auth.RegisterAuthServiceServer(grpcServer, &AuthServer{})
	user.RegisterUserServiceServer(grpcServer, &UserServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server on API Gateway Service: %v", err)
	}
	defer listener.Close()
}
