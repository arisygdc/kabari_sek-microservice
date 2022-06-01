package main

import (
	"chat-in-app_microservices/api-gateway/clientapi"
	"chat-in-app_microservices/api-gateway/config"
	"chat-in-app_microservices/api-gateway/router"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Config{}
	err := config.LoadConfig(".", "config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	engine := mux.NewRouter()
	svcClient := clientapi.NewService(cfg.Endpoint)
	router := router.NewRouter(cfg.Server, engine, svcClient)
	router.RegisterRoute()

	err = router.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
