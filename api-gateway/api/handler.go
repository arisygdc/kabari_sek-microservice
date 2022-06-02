package api

import (
	"chat-in-app_microservices/api-gateway/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// Handler server using gorilla mux
type Handler struct {
	cfg    config.ConfigServer
	engine *mux.Router
}

func NewHandler(cfg config.ConfigServer, handler *mux.Router) Handler {
	return Handler{
		cfg:    cfg,
		engine: handler,
	}
}

func (h Handler) Post(path string, f HandlerFunc) {
	h.engine.HandleFunc(path, f).Methods(http.MethodPost)
}

func (h Handler) Get(path string, f HandlerFunc) {
	h.engine.HandleFunc(path, f).Methods(http.MethodGet)
}

func (h Handler) Serve() error {
	addr := fmt.Sprintf("%s:%d", h.cfg.Host, h.cfg.Port)
	log.Printf("server running on : %s", addr)
	h.engine.NotFoundHandler = http.NotFoundHandler()
	return http.ListenAndServe(addr, h.engine)
}
