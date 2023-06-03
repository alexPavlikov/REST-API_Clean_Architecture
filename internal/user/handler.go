package user

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
	userURL  = "/user/:uuid"
	usersURL = "/users"
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
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUser))
	router.HandlerFunc(http.MethodPost, userURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.GetAll(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	usBytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(usBytes)
	return nil
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	user, err := h.service.GetOne(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	usBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usBytes)
	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	user := User{
		Id:           "",
		Firstname:    "",
		Lastname:     "",
		Age:          0,
		Email:        "",
		PasswordHash: "",
	}
	err := h.service.Create(context.TODO(), &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create user complited"))
	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	user := User{
		Id:           "",
		Firstname:    "",
		Lastname:     "",
		Age:          0,
		Email:        "",
		PasswordHash: "",
	}
	err := h.service.Update(context.TODO(), &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	usBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usBytes)
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	err := h.service.Delete(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Delete user by id: %s", id)))
	return nil
}
