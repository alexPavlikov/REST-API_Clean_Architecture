package worker

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/apperror"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/handlers"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	workerURL  = "/worker/:uuid"
	workersURL = "/workers"
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

func (h *handler) Register(route *httprouter.Router) {
	route.HandlerFunc(http.MethodGet, workersURL, apperror.Middleware(h.GetList))
	route.HandlerFunc(http.MethodGet, workerURL, apperror.Middleware(h.GetOne))
	route.HandlerFunc(http.MethodPost, workerURL, apperror.Middleware(h.CreateWorker))
	route.HandlerFunc(http.MethodPatch, workerURL, apperror.Middleware(h.UpdateWorker))
	route.HandlerFunc(http.MethodDelete, workerURL, apperror.Middleware(h.DeleteWorker))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	workers, err := h.service.GetAll(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	workBytes, err := json.Marshal(workers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(workBytes)
	return nil
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/worker/")
	worker, err := h.service.GetOne(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	wBytes, err := json.Marshal(worker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(wBytes)
	return nil
}

func (h *handler) CreateWorker(w http.ResponseWriter, r *http.Request) error {
	worker := Workrer{
		Id:           "",
		Firstname:    "",
		Lastname:     "",
		Age:          0,
		Experieons:   0,
		Number:       "",
		Address:      "",
		Email:        "",
		PasswordHash: "",
	}
	err := h.service.CreateWorker(context.TODO(), &worker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("create worker - complited"))
	return nil
}

func (h *handler) UpdateWorker(w http.ResponseWriter, r *http.Request) error {
	worker := Workrer{
		Id:           "",
		Firstname:    "",
		Lastname:     "",
		Age:          0,
		Experieons:   0,
		Number:       "",
		Address:      "",
		Email:        "",
		PasswordHash: "",
	}
	err := h.service.UpdateWorker(context.TODO(), &worker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("update worker - complited"))
	return nil
}

func (h *handler) DeleteWorker(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/worker/")
	err := h.service.DeleteWorker(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted worker - complited"))
	return nil
}
