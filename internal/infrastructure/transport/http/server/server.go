package server

import (
	"avitoTestTask/internal/config"
	"avitoTestTask/internal/infrastructure/transport/http/handlers"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
	cfg        *config.Server
}

func NewServer(h *handlers.Handlers, cfg *config.Server) *Server {
	router := initRoutes(h)

	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			Handler:      router,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	serverErrors := make(chan error, 1)

	go func() {
		logrus.Infof("http-server is starting on port %d...", s.cfg.Port)
		serverErrors <- s.httpServer.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return err
	case sig := <-quit:
		logrus.Infof("Received shutdown signal: %v", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.httpServer.Shutdown(ctx); err != nil {
			logrus.Errorf("Graceful shutdown failed: %v", err)
			return s.httpServer.Close()
		}
		logrus.Info("Server shutdown completed")
		return nil
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
