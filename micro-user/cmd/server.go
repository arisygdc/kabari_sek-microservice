package main

import (
	"chat-in-app_microservices/micro-user/config"
	"chat-in-app_microservices/micro-user/core"
	"chat-in-app_microservices/micro-user/core/user"
	"chat-in-app_microservices/micro-user/pb"
	"log"

	"go-micro.dev/v4"
)

func main() {
	svcServer := micro.NewService(
		micro.Name("svc-user"),
	)

	conf := config.Config{}
	err := config.LoadConfig("../", "config", &conf)
	if err != nil {
		log.Fatal(err)
	}

	repos, err := core.NewRepository(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	svcUser := user.NewService(repos)
	router := user.NewUserRouter(conf, svcUser)
	err = pb.RegisterUserHandler(svcServer.Server(), router)

	if err != nil {
		log.Fatal(err)
	}

	svcServer.Init()
}
