package app

import (
	"avitoTestTask/internal/config"
	"avitoTestTask/internal/infrastructure/repository/postgres"
	"avitoTestTask/internal/usecases"

	"github.com/sirupsen/logrus"
)

func RunServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}
	logrus.Info(cfg)

	repo := postgres.NewPostgresRepo()
	usecases := usecases.NewUseCases(repo.UserRepo, repo.TeamRepo, repo.PullRequestRepo)
	_ = usecases //REMOVE
}
