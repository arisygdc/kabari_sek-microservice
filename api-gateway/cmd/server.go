package main

import (
	"chat-in-app_microservices/api-gateway/clientapi"
	"chat-in-app_microservices/api-gateway/router"

	"github.com/gorilla/mux"
)

func main() {
	engine := mux.NewRouter()
	router := router.NewRouter(engine, clientapi.NewService())
	router.RegisterRoute()
	router.Serve()
}
