package main

import (
	"chat-in-app_microservices/api-gateway/config"
	consumerUserSvc "chat-in-app_microservices/api-gateway/core/api/grpc"
	httpServer "chat-in-app_microservices/api-gateway/core/api/http"
	"log"
)

func main() {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	svcClient := consumerUserSvc.NewService(cfg.Endpoint)
	srv := httpServer.NewEchoServer(cfg, svcClient)
	err = srv.RegisterRoute()

	if err != nil {
		log.Fatal(err)
	}
}
