package postgres

import "gorm.io/gorm"

type PostgresRepo struct {
	UserRepo        *UserRepo
	TeamRepo        *TeamRepo
	PullRequestRepo *PullRequestRepo
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		UserRepo:        NewUserRepo(db),
		TeamRepo:        NewTeamRepo(db),
		PullRequestRepo: NewPullRequestRepo(db),
	}
}
