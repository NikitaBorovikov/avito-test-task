package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
)

type PullRequestUC struct {
	PullRequestRepo repository.PullRequestRepo
}

func NewPullRequestUC(pullRequestRepo repository.PullRequestRepo) *PullRequestUC {
	return &PullRequestUC{
		PullRequestRepo: pullRequestRepo,
	}
}

func (uc *PullRequestUC) Create(pr *models.PullRequest) (*models.PullRequest, error) {
	return uc.PullRequestRepo.Create(pr)
}

func (uc *PullRequestUC) GetByReviewer(userID string) ([]models.PullRequest, error) {
	return uc.PullRequestRepo.GetByReviewer(userID)
}

func (uc *PullRequestUC) Merge(prID string) (*models.PullRequest, error) {
	return uc.PullRequestRepo.Merge(prID)
}

func (uc *PullRequestUC) Reassign(prID, oldUserID string) (*models.PullRequest, error) {
	return nil, nil
}
