package http

import (
	"chat-in-app_microservices/api-gateway/core/api/grpc"
	"chat-in-app_microservices/api-gateway/pkg/pb"
	"fmt"

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

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birth     string `json:"birth"`
}

func (h Handler) Register(ctx echo.Context) (err error) {
	req := new(RegisterRequest)
	if err = ctx.Bind(req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	usernameCreated, err := h.svc.User.Register(ctx.Request().Context(), &pb.RegisterRequest{
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
		"message": fmt.Sprintf("%s created", usernameCreated),
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h Handler) Login(ctx echo.Context) (err error) {
	req := new(LoginRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	_, err = h.svc.User.Login(ctx.Request().Context(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	return
}
