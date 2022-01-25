package common

import (
	"os"
)

type ConfigStore struct {
	config *Config
}

type Config struct {
	GrpcPort    string
	GrpcAddress string
	RestPort    string
}

func GetConfig() *Config {

	GRPC_SERVER_PORT := os.Getenv("NET_REPLY_SERVICE_SERVICE_PORT_GRPC")
	GRCP_SERVER_NETWORK_ADDRESS := os.Getenv("NET_REPLY_SERVICE_SERVICE_HOST")
	REST_SERVER_PORT := os.Getenv("NET_REPLY_SERVICE_SERVICE_PORT_HTTP")

	if GRPC_SERVER_PORT == "" {
		GRPC_SERVER_PORT = "8080"
	}
	if GRCP_SERVER_NETWORK_ADDRESS == "" {
		GRCP_SERVER_NETWORK_ADDRESS = "localhost"
	}
	if REST_SERVER_PORT == "" {
		REST_SERVER_PORT = "8081"
	}
	return &Config{
		GrpcPort:    GRPC_SERVER_PORT,
		GrpcAddress: GRCP_SERVER_NETWORK_ADDRESS,
		RestPort:    REST_SERVER_PORT,
	}
}
