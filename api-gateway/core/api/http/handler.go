package http

import (
	"chat-in-app_microservices/api-gateway/core/api/grpc"
	"chat-in-app_microservices/api-gateway/pkg/pb"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc grpc.ServiceGrpcAPI
}

func NewHandler(svc grpc.ServiceGrpcAPI) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) Register(ctx echo.Context) (err error) {
	req := new(RegisterRequest)
	if err = ctx.Bind(req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respSvc, err := h.svc.User.Register(ctx.Request().Context(), &pb.RegisterRequest{
		Username:  req.Username,
		Password:  req.Password,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Birth:     req.Birth,
	})

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(201, map[string]interface{}{
		"message": fmt.Sprintf("%s created", respSvc.Username),
	})
}

func (h Handler) Login(ctx echo.Context) (err error) {
	req := new(LoginRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := ValidateRequest(req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	respSvc, err := h.svc.User.Login(ctx.Request().Context(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "X-SESSION-TOKEN",
		Value:    respSvc.UserToken,
		HttpOnly: true,
	})

	return ctx.JSON(200, map[string]interface{}{
		"message": "logged in",
	})
}
