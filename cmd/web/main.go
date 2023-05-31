package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/book"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/config"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/user---"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/client/postgresql"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	repBook := book.NewRepository(postgreSQLClient, logger)

	books, err := repBook.FindAll(context.TODO())
	if err != nil {
		logger.Fatalf("%v", err)
	}

	for _, b := range books {
		logger.Infof("%v", b)
	}

	logger.Info("Register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, config *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Start application")
	var listener net.Listener
	var listenErr error

	logger.Info("Listen TCP")
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", config.Listen.BindIP, config.Listen.Port))
	logger.Infof("Server is listening port %s:%s", config.Listen.BindIP, config.Listen.Port)
	if listenErr != nil {
		logger.Fatal(listenErr.Error())
	}

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	err := server.Serve(listener)
	if err != nil {
		logger.Fatal(err.Error())
	}

}
