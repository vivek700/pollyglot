package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/openai/openai-go/v3"
)

type Server struct {
	port    int
	logger  *slog.Logger
	client  openai.Client
	aiModel string
}

func NewServer(logger *slog.Logger, port int, client openai.Client, aiModel string) *http.Server {
	newServer := &Server{
		port:    port,
		logger:  logger,
		client:  client,
		aiModel: aiModel,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}
