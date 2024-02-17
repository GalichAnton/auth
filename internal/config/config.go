package config

import "github.com/joho/godotenv"

// Load - .
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPC - .
type GRPC interface {
	Address() string
}

// PGConfig - .
type PGConfig interface {
	DSN() string
}
