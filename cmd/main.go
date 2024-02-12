package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/GalichAnton/auth/cmd/server"
	"github.com/GalichAnton/auth/internal/repository/pg"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

const grpcPort = 50051
const (
	dbDSN = "host=localhost port=54321 dbname=users user=anton password=admin sslmode=disable"
)

func main() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepository := pg.NewUserRepository(pool)

	s := grpc.NewServer()
	reflection.Register(s)

	userServer := server.NewUserServer(userRepository)

	desc.RegisterUserV1Server(s, userServer)

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
