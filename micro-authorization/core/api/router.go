package api

import (
	"chat-in-app_microservices/micro-authorization/core/authorization"
	"chat-in-app_microservices/micro-authorization/pb"
	"context"
	"net/http"

	"github.com/google/uuid"
)

type GrpcHandler struct {
	svc authorization.AuthorizationService
}

func NewGrpcHandler(svc authorization.AuthorizationService) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (rpc GrpcHandler) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest, res *pb.CheckPermissionResponse) error {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return err
	}

	svcPermission, err := rpc.svc.UserGetPermission(ctx, id, req.PermissionName)
	if err != nil {
		return err
	}

	res.ResponseCode = http.StatusFound
	res.Message = svcPermission.Name
	return nil
}

func (rpc GrpcHandler) BannStatus(ctx context.Context, req *pb.BannStatusRequest, res *pb.BannStatusResponse) error {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return err
	}

	_, err = rpc.svc.UserGetBannedPermission(ctx, id, req.PermissionName)
	if err != nil {
		return err
	}

	res.ResponseCode = http.StatusFound
	res.Message = req.PermissionName
	return nil
}
