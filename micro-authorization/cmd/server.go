package main

import (
	"chat-in-app_microservices/micro-authorization/config"
	"chat-in-app_microservices/micro-authorization/core/api"
	"chat-in-app_microservices/micro-authorization/core/authorization"
	"chat-in-app_microservices/micro-authorization/db"
	"chat-in-app_microservices/micro-authorization/pb"
	"context"
	"log"

	"go-micro.dev/v4"
)

func main() {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	svcServer := micro.NewService(
		micro.Name(cfg.Service.Name),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repos, err := db.NewRepository(ctx, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	svcAuthorization := authorization.NewAuthorization(repos)
	grpcHandler := api.NewGrpcHandler(svcAuthorization)

	err = pb.RegisterAuthorizationHandler(svcServer.Server(), grpcHandler)
	if err != nil {
		log.Fatal(err)
	}
}
