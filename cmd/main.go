package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/GalichAnton/auth/cmd/server"
	"github.com/GalichAnton/auth/internal/config"
	"github.com/GalichAnton/auth/internal/config/env"
	"github.com/GalichAnton/auth/internal/repository/pg"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to parse gRPC config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to parse PG config: %v", err)
	}

	ctx := context.Background()
	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
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
