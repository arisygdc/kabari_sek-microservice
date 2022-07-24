package main

import (
	"chat-in-app_microservices/api-gateway/clientapi"
	"chat-in-app_microservices/api-gateway/config"
	"log"
)

func main() {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	clientapi.NewService(cfg.Endpoint)

	if err != nil {
		log.Fatal(err)
	}
}
