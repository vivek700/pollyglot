package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port   int
	logger *slog.Logger
}

func NewServer(logger *slog.Logger) *http.Server {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Error("Invalid or missing port", "err", err)
		os.Exit(1)
	}

	newServer := &Server{
		port:   port,
		logger: logger,
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
