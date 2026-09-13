package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/pollyglot/internal/server"
)

func main() {

	logger := initLogger()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Error("Invalid or missing port", "err", err)
		os.Exit(1)
	}

	apiKey := os.Getenv("AI_KEY")
	baseUrl := os.Getenv("AI_URL")
	aiModel := os.Getenv("AI_MODEL")

	if apiKey == "" || baseUrl == "" || aiModel == "" {
		logger.Error("Missing api key or base url or ai model")
		os.Exit(1)
	}

	client := openai.NewClient(option.WithAPIKey(apiKey), option.WithBaseURL(baseUrl))

	srv := server.NewServer(logger, port, client, aiModel)

	done := make(chan struct{})

	go gracefulShutdown(srv, logger, done)

	logger.Debug("Server is starting", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("http server error", "err", err)
		os.Exit(1)
	}

	<-done
	logger.Info("Graceful shutdown complete.")

}

func initLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

}

func gracefulShutdown(apiServer *http.Server, logger *slog.Logger, done chan struct{}) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	logger.Info("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown with error", "err", err)
	}

	logger.Info("Server exiting")

	close(done)

}
