package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type IRoute interface {
	Post(path string, f HandlerFunc)
	Get(path string, f HandlerFunc)
}

// Handler server using gorilla mux
type Handler struct {
	r *mux.Router
}

func NewHandler(handler *mux.Router) Handler {
	return Handler{
		r: handler,
	}
}

func (h Handler) Post(path string, f HandlerFunc) {
	h.r.HandleFunc(path, f).Methods(http.MethodPost)
}

func (h Handler) Get(path string, f HandlerFunc) {
	h.r.HandleFunc(path, f).Methods(http.MethodGet)
}
