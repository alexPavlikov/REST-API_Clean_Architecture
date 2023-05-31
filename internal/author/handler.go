package author

import (
	"net/http"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/handlers"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	authorURL  = "/author/:uuid"
	authorsURL = "/authors"
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
	router.HandlerFunc(http.MethodGet, authorsURL, h.GetList)
	router.HandlerFunc(http.MethodGet, authorURL, h.GetAuthor)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is list all authors"))
}

func (h *handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is page of author"))
}
