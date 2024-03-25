package env

import (
	"errors"
	"net"
	"os"
)

const (
	prometheusHost = "PROMETHEUS_HOST"
	prometheusPort = "PROMETHEUS_PORT"
)

// PrometheusConfig ...
type PrometheusConfig interface {
	Address() string
}

type prometheusConfig struct {
	host string
	port string
}

// NewPrometheusConfig ...
func NewPrometheusConfig() (PrometheusConfig, error) {
	host := os.Getenv(prometheusHost)
	if len(host) == 0 {
		return nil, errors.New("prometheus host not found")
	}

	port := os.Getenv(prometheusPort)
	if len(port) == 0 {
		return nil, errors.New("prometheus port not found")
	}

	return &prometheusConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *prometheusConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
