package api

import (
	"chat-in-app_microservices/api-gateway/clientapi"
	"chat-in-app_microservices/api-gateway/config"
	"chat-in-app_microservices/api-gateway/pb"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	CookieSessionName = "session"
)

type Router struct {
	cfg     config.ConfigServer
	engine  *mux.Router
	service clientapi.ServiceAPI
	helper  IHttpHelper
	route   IRoute
}

//
func NewRouter(cfg config.ConfigServer, handler *mux.Router, svc clientapi.ServiceAPI) Router {
	return Router{
		cfg:     cfg,
		engine:  handler,
		service: svc,
		helper:  NewHttpHelper(),
		route:   NewHandler(handler),
	}
}

// All route goes here
func (r Router) RegisterRoute() {
	r.route.Post("/login", r.login)
	r.route.Post("/register", r.register)
}

func (r Router) Serve() error {
	addr := fmt.Sprintf("%s:%d", r.cfg.Host, r.cfg.Port)
	return http.ListenAndServe(addr, r.engine)
}

func (router Router) login(w http.ResponseWriter, r *http.Request) {
	var req LoginPost
	if err := router.helper.JsonParse(r.Body, req); err != nil {
		return
	}

	loginResponse, err := router.service.User.Login(r.Context(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		router.helper.JsonResponseMessage(w, int(loginResponse.GetResponseCode()), loginResponse.GetMessage())
		return
	}

	responseCode := int(loginResponse.GetResponseCode())
	if responseCode != http.StatusAccepted {
		router.helper.JsonResponseMessage(w, responseCode, loginResponse.GetMessage())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CookieSessionName,
		Value:    loginResponse.GetUserToken(),
		HttpOnly: true,
		Secure:   true,
	})

	router.helper.JsonResponseMessage(w, http.StatusAccepted, "login success")
}

func (router Router) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterPost
	if err := router.helper.JsonParse(r.Body, req); err != nil {
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
		router.helper.JsonResponseMessage(w, int(registerResponse.GetResponseCode()), registerResponse.GetMessage())
		return
	}

	responseCode := int(registerResponse.GetResponseCode())
	if responseCode != http.StatusAccepted {
		router.helper.JsonResponseMessage(w, responseCode, registerResponse.GetMessage())
		return
	}

	router.helper.JsonResponseMessage(w, http.StatusCreated, http.StatusText(http.StatusCreated))
}