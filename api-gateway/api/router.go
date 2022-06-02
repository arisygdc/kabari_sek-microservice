package api

import (
	"chat-in-app_microservices/api-gateway/clientapi"
	"chat-in-app_microservices/api-gateway/config"
	"chat-in-app_microservices/api-gateway/pb"
	"net/http"
)

type IRoute interface {
	Post(path string, f HandlerFunc)
	Get(path string, f HandlerFunc)
}

var (
	CookieSessionName = "session"
)

type Router struct {
	service clientapi.ServiceGrpcAPI
	helper  IHttpHelper
}

func NewRouter(cfg config.ConfigServer, svc clientapi.ServiceGrpcAPI) Router {
	return Router{
		service: svc,
		helper:  NewHttpHelper(),
	}
}

// All route goes here
func (r Router) RegisterRoute(routeHandler IRoute) {
	routeHandler.Post("/login", r.login)
	routeHandler.Post("/register", r.register)
}

func (router Router) login(w http.ResponseWriter, r *http.Request) {
	var req LoginPost
	if err := router.helper.JsonParse(r.Body, &req); err != nil {
		return
	}

	pbLoginReq := pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	loginResponse, err := router.service.User.Login(r.Context(), &pbLoginReq)

	if err != nil {
		if loginResponse == nil {
			router.helper.JsonResponseMessage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		router.helper.JsonResponseMessage(w, int(loginResponse.ResponseCode), loginResponse.Message)
		return
	}

	responseCode := int(loginResponse.ResponseCode)
	if responseCode != http.StatusAccepted {
		router.helper.JsonResponseMessage(w, responseCode, loginResponse.Message)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CookieSessionName,
		Value:    loginResponse.UserToken,
		HttpOnly: true,
		Secure:   true,
	})

	router.helper.JsonResponseMessage(w, http.StatusAccepted, "login success")
}

func (router Router) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterPost
	if err := router.helper.JsonParse(r.Body, &req); err != nil {
		router.helper.JsonResponseMessage(w, http.StatusBadRequest, err)
		return
	}

	registerResponse, err := router.service.User.Register(r.Context(), &pb.RegisterRequest{
		Username:  req.Usename,
		Password:  req.Password,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Birth:     req.Birth,
	})

	if err != nil {
		if registerResponse == nil {
			router.helper.JsonResponseMessage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		router.helper.JsonResponseMessage(w, int(registerResponse.ResponseCode), registerResponse.Message)
		return
	}

	responseCode := int(registerResponse.ResponseCode)
	if responseCode != http.StatusAccepted {
		router.helper.JsonResponseMessage(w, responseCode, registerResponse.Message)
		return
	}

	router.helper.JsonResponseMessage(w, http.StatusCreated, http.StatusText(http.StatusCreated))
}
