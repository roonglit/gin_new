package config

type Config struct {
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}
