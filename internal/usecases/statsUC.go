package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
)

type StatsUC struct {
	StatsRepo repository.StatsRepo
}

func NewStatsUC(statsRepo repository.StatsRepo) *StatsUC {
	return &StatsUC{
		StatsRepo: statsRepo,
	}
}

func (uc *StatsUC) GetReviewerStats() ([]models.ReviewerStats, error) {
	return uc.StatsRepo.GetReviewerStats()
}
