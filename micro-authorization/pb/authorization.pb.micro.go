// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authorization.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Authorization service

func NewAuthorizationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Authorization service

type AuthorizationService interface {
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...client.CallOption) (*CheckPermissionResponse, error)
	BannStatus(ctx context.Context, in *BannStatusRequest, opts ...client.CallOption) (*BannStatusResponse, error)
	GiveRole(ctx context.Context, in *GiveRoleRequest, opts ...client.CallOption) (*GiveRoleResponse, error)
}

type authorizationService struct {
	c    client.Client
	name string
}

func NewAuthorizationService(name string, c client.Client) AuthorizationService {
	return &authorizationService{
		c:    c,
		name: name,
	}
}

func (c *authorizationService) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...client.CallOption) (*CheckPermissionResponse, error) {
	req := c.c.NewRequest(c.name, "Authorization.CheckPermission", in)
	out := new(CheckPermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationService) BannStatus(ctx context.Context, in *BannStatusRequest, opts ...client.CallOption) (*BannStatusResponse, error) {
	req := c.c.NewRequest(c.name, "Authorization.BannStatus", in)
	out := new(BannStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationService) GiveRole(ctx context.Context, in *GiveRoleRequest, opts ...client.CallOption) (*GiveRoleResponse, error) {
	req := c.c.NewRequest(c.name, "Authorization.GiveRole", in)
	out := new(GiveRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationHandler interface {
	CheckPermission(context.Context, *CheckPermissionRequest, *CheckPermissionResponse) error
	BannStatus(context.Context, *BannStatusRequest, *BannStatusResponse) error
	GiveRole(context.Context, *GiveRoleRequest, *GiveRoleResponse) error
}

func RegisterAuthorizationHandler(s server.Server, hdlr AuthorizationHandler, opts ...server.HandlerOption) error {
	type authorization interface {
		CheckPermission(ctx context.Context, in *CheckPermissionRequest, out *CheckPermissionResponse) error
		BannStatus(ctx context.Context, in *BannStatusRequest, out *BannStatusResponse) error
		GiveRole(ctx context.Context, in *GiveRoleRequest, out *GiveRoleResponse) error
	}
	type Authorization struct {
		authorization
	}
	h := &authorizationHandler{hdlr}
	return s.Handle(s.NewHandler(&Authorization{h}, opts...))
}

type authorizationHandler struct {
	AuthorizationHandler
}

func (h *authorizationHandler) CheckPermission(ctx context.Context, in *CheckPermissionRequest, out *CheckPermissionResponse) error {
	return h.AuthorizationHandler.CheckPermission(ctx, in, out)
}

func (h *authorizationHandler) BannStatus(ctx context.Context, in *BannStatusRequest, out *BannStatusResponse) error {
	return h.AuthorizationHandler.BannStatus(ctx, in, out)
}

func (h *authorizationHandler) GiveRole(ctx context.Context, in *GiveRoleRequest, out *GiveRoleResponse) error {
	return h.AuthorizationHandler.GiveRole(ctx, in, out)
}