package env

import (
	"errors"
	"net"
	"os"

	"github.com/GalichAnton/auth/internal/config"
)

var _ config.GRPC = (*GRPCConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc port not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &GRPCConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
