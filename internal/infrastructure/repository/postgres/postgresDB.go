package postgres

import (
	"strings"

	"gorm.io/gorm"
)

const (
	PostgresUniqueErrorCode = "23505"
)

type PostgresRepo struct {
	UserRepo        *UserRepo
	TeamRepo        *TeamRepo
	PullRequestRepo *PullRequestRepo
	StatsRepo       *StatsRepo
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		UserRepo:        NewUserRepo(db),
		TeamRepo:        NewTeamRepo(db),
		PullRequestRepo: NewPullRequestRepo(db),
		StatsRepo:       NewStatsRepo(db),
	}
}

func isDuplicateError(err error) bool {
	return strings.Contains(err.Error(), PostgresUniqueErrorCode)
}
