package main

import (
	"chat-in-app_microservices/api-gateway/api"
	"chat-in-app_microservices/api-gateway/clientapi"
	"chat-in-app_microservices/api-gateway/config"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	svcClient := clientapi.NewService(cfg.Endpoint)

	engine := mux.NewRouter()
	handler := api.NewHandler(cfg.Server, engine)

	router := api.NewRouter(cfg.Server, svcClient)
	router.RegisterRoute(handler)

	err = handler.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
