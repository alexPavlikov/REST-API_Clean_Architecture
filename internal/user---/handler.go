package user

import (
	"net/http"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/handlers"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	userURL  = "/user/:uuid"
	usersURL = "/users"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handlers {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, h.GetList)
	router.HandlerFunc(http.MethodGet, userURL, h.GetUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is list all users"))
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is page of user"))
}
