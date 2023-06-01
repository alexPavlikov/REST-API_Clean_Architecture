package author

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/apperror"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/handlers"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	authorURL  = "/author/:uuid"
	authorsURL = "/authors"
)

type handler struct {
	logger  *logging.Logger
	service *Service
}

func NewHandler(logger *logging.Logger, service *Service) handlers.Handlers {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, authorURL, apperror.Middleware(h.GetAuthor))
	router.HandlerFunc(http.MethodPost, authorURL, apperror.Middleware(h.CreateAuthor))
	router.HandlerFunc(http.MethodPatch, authorURL, apperror.Middleware(h.UpdateAuthor))
	router.HandlerFunc(http.MethodDelete, authorURL, apperror.Middleware(h.DeleteAuthor))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.service.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
	return nil
}

func (h *handler) GetAuthor(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/author/")
	ath, err := h.service.GetOne(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(200)
	w.Write([]byte(ath.ID + " " + ath.Name))
	return nil
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) error {
	name := strings.TrimPrefix(r.URL.Path, "/authors/")
	err := h.service.CreateAuthor(context.TODO(), name)
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("create new author - %s", name)))
	return nil
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) error {
	ath := Author{
		ID:   "",
		Name: "",
	}
	err := h.service.UpdateAuthour(context.TODO(), ath)
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("update author - ID: %s; name: %s", ath.ID, ath.Name)))
	return nil
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/authors/")
	err := h.service.DeleteAuthor(context.TODO(), id)
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("deleted author by ID: %s", id)))
	return nil
}
