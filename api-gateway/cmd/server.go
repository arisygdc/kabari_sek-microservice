package main

import (
	"chat-in-app_microservices/api-gateway/config"
	consumerUserSvc "chat-in-app_microservices/api-gateway/core/api/grpc"
	"log"
)

func main() {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	consumerUserSvc.NewService(cfg.Endpoint)

	if err != nil {
		log.Fatal(err)
	}
}
