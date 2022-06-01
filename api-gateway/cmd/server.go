package main

import (
	"chat-in-app_microservices/api-gateway/pb"
	"chat-in-app_microservices/api-gateway/router"

	"github.com/gorilla/mux"
	"go-micro.dev/v4/client"
)

func main() {
	engine := mux.NewRouter()
	router := router.NewRouter(engine, pb.NewUserService("svc_user", client.DefaultClient))
	router.RegisterRoute()
	router.Serve()
}
