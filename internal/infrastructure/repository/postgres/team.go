package postgres

import (
	"avitoTestTask/internal/core/models"

	"gorm.io/gorm"
)

type TeamRepo struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{
		db: db,
	}
}

func (r *TeamRepo) Create(team *models.Team) (*models.Team, error) {
	return nil, nil
}

func (r *TeamRepo) GetByName(name string) (*models.Team, error) {
	return nil, nil
}

func (r *TeamRepo) Delete(teamID uint) error {
	return nil
}
