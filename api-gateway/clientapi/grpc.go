package clientapi

import (
	"chat-in-app_microservices/api-gateway/pb"

	"chat-in-app_microservices/api-gateway/config"

	"go-micro.dev/v4/client"
	"golang.org/x/net/context"
)

// Client GRPC api
type ServiceGrpcAPI struct {
	User pb.UserService
}

func NewService(cfg config.ConfigServiceEndpoint) ServiceGrpcAPI {
	rpcClient := client.NewClient(func(o *client.Options) {
		o.Context = context.Background()
	})

	return ServiceGrpcAPI{
		User: pb.NewUserService(cfg.User, rpcClient),
	}
}
