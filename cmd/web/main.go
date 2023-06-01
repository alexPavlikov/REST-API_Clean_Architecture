package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/author"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/config"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/client/postgresql"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	clientPostgreSQL, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	logger.Info("Register authors handler")
	authorRepos := author.NewRepository(clientPostgreSQL, logger)
	authorService := author.NewService(authorRepos, logger)
	authorsHandler := author.NewHandler(logger, authorService)
	authorsHandler.Register(router)

	// logger.Info("Register user handler")
	// userHandler := user.NewHandler(logger)
	// userHandler.Register(router)

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
