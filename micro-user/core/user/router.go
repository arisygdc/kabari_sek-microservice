package user

import (
	"chat-in-app_microservices/micro-user/config"
	"chat-in-app_microservices/micro-user/pkg/pb"
	"context"
	"net/http"
)

type UserRouter struct {
	conf config.Config
	svc  Service
}

func NewUserRouter(conf config.Config, svc Service) UserRouter {
	return UserRouter{
		conf: conf,
		svc:  svc,
	}
}

func (r UserRouter) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {
	loginParam := Auth{username: req.Username, password: req.Password}
	token, err := r.svc.Login(ctx, loginParam, r.conf.Token.Secret, r.conf.Token.Dur)
	if err != nil {
		if err == ErrIncorrectUserOrPass {
			res.ResponseCode = http.StatusForbidden
			res.Message = err.Error()
		}

		res.ResponseCode = http.StatusInternalServerError
		res.Message = http.StatusText(int(res.ResponseCode))
		return err
	}

	res.ResponseCode = http.StatusAccepted
	res.Message = http.StatusText(int(res.ResponseCode))
	res.UserToken = token
	return nil
}

func (r UserRouter) Register(ctx context.Context, req *pb.RegisterRequest, res *pb.RegisterResponse) error {
	birth, err := ParseTime(req.Birth)
	if err != nil {
		res.ResponseCode = http.StatusBadRequest
		res.Message = err.Error()
		return err
	}

	registerParam := RegisterUser{
		Auth: Auth{
			username: req.Username,
			password: req.Password,
		},
		User: User{
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Birth:     birth,
		},
	}

	err = r.svc.Register(ctx, registerParam)
	if err != nil {
		res.ResponseCode = http.StatusInternalServerError
		res.Message = http.StatusText(int(res.ResponseCode))
		return err
	}

	res.ResponseCode = http.StatusAccepted
	res.Message = http.StatusText(int(res.ResponseCode))
	return nil
}
