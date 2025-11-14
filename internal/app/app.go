package app

import (
	"avitoTestTask/internal/config"
	"avitoTestTask/internal/infrastructure/repository/postgres"
	"avitoTestTask/internal/infrastructure/transport/http/handlers"
	"avitoTestTask/internal/infrastructure/transport/http/server"
	"avitoTestTask/internal/usecases"

	"github.com/sirupsen/logrus"
)

func RunServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	repo := postgres.NewPostgresRepo()
	usecases := usecases.NewUseCases(repo.UserRepo, repo.TeamRepo, repo.PullRequestRepo)
	handlers := handlers.NewHandlers(usecases)

	httpServer := server.NewServer(handlers, &cfg.Server)
	if err := httpServer.Run(); err != nil {
		logrus.Fatalf("server error: %v", err)
	}
}
