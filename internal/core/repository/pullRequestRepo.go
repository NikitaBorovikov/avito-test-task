package repository

import "avitoTestTask/internal/core/models"

type PullRequestRepo interface {
	Create(pr *models.PullRequest) (*models.PullRequest, error)
	Update(pr *models.PullRequest) (*models.PullRequest, error)
	GetByReviewer(userID string) ([]models.PullRequest, error)
	Merge(prID string) (*models.PullRequest, error)
	Reassign(prID, oldUserID string) (*models.PullRequest, error)
}
