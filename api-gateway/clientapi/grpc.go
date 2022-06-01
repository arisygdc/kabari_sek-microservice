package clientapi

import (
	"chat-in-app_microservices/api-gateway/pb"

	"go-micro.dev/v4/client"
	"golang.org/x/net/context"
)

const (
	ServiceUserEndpoint string = "svc-user"
)

type ServiceAPI struct {
	User pb.UserService
}

func NewService() ServiceAPI {
	rpcClient := client.NewClient(func(o *client.Options) {
		o.Context = context.Background()
	})

	return ServiceAPI{
		User: pb.NewUserService(ServiceUserEndpoint, rpcClient),
	}
}
