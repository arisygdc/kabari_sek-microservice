package main

import (
	"chat-in-app_microservices/api-gateway/config"
	consumerUserSvc "chat-in-app_microservices/api-gateway/core/api/grpc"
	httpServer "chat-in-app_microservices/api-gateway/core/api/http"
	"log"
)

func main() {
	srv := serverInit()
	srv.RegisterRoute()
}

func serverInit() (srv httpServer.IHttpServer) {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	svcClient := consumerUserSvc.NewService(cfg.Endpoint)
	srv = httpServer.NewEchoServer(cfg, svcClient)
	return
}
