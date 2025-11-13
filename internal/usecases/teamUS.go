package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
)

type TeamUC struct {
	TeamRepo repository.TeamRepo
}

func NewTeamUC(teamRepo repository.TeamRepo) *TeamUC {
	return &TeamUC{
		TeamRepo: teamRepo,
	}
}

func (uc *TeamUC) Create(team *models.Team) (*models.Team, error) {
	return nil, nil
}

func (uc *TeamUC) GetByName(name string) (*models.Team, error) {
	return nil, nil
}
