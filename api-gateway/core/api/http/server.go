package http

import (
	"chat-in-app_microservices/api-gateway/config"
	"chat-in-app_microservices/api-gateway/core/api/grpc"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IHttpServer interface {
	RegisterRoute()
}

type EchoServer struct {
	cfg      config.Config
	handler  Handler
	provider *echo.Echo
}

func NewEchoServer(cfg config.Config, svc grpc.ServiceGrpcAPI) EchoServer {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return EchoServer{
		cfg:      cfg,
		handler:  NewHandler(svc),
		provider: e,
	}
}

func (e EchoServer) RegisterRoute() {
	e.provider.GET("/api/v1/register", e.handler.Register)
}
