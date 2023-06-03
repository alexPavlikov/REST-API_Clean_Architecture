package book

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/apperror"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/author"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/handlers"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	book      = "/book/:uuid"
	books     = "/books"
	booksUUID = "/books/:uuid"
)

type handler struct {
	service *Service
	logger  *logging.Logger
}

// Register implements handlers.Handlers.
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, books, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, book, apperror.Middleware(h.GetOne))
	router.HandlerFunc(http.MethodPost, booksUUID, apperror.Middleware(h.GetAllByAuthor))
	router.HandlerFunc(http.MethodPost, book, apperror.Middleware(h.CreateBook))
	router.HandlerFunc(http.MethodPatch, book, apperror.Middleware(h.UpdateBook))
	router.HandlerFunc(http.MethodDelete, book, apperror.Middleware(h.DeleteBook))
}

func NewHandler(service *Service, logger *logging.Logger) handlers.Handlers {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	books, err := h.service.GetAll(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	booksBytes, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(booksBytes)
	return nil
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/book/")
	fmt.Println(id)
	book, err := h.service.GetOne(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	bBytes, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bBytes)
	return nil
}

func (h *handler) GetAllByAuthor(w http.ResponseWriter, r *http.Request) error {

	id := strings.TrimPrefix(r.URL.Path, "/books/")

	author, err := h.service.GetAllByAuthor(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	athrBytes, err := json.Marshal(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(athrBytes)
	return nil
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) error {
	book := Book{
		ID:   "",
		Name: "",
		Author: []author.Author{
			{
				ID:   "",
				Name: "",
			},
		},
	}
	err := h.service.CreateBook(context.TODO(), &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) error {
	book := Book{
		ID:   "",
		Name: "",
		Author: []author.Author{{
			ID:   "",
			Name: "",
		}},
	}
	err := h.service.UpdateBook(context.TODO(), &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/book/")
	err := h.service.DeleteBook(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
