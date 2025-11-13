package repository

import "avitoTestTask/internal/core/models"

type PRRepo interface {
	Create(pr *models.PullRequest) (*models.PullRequest, error)
	Update(pr *models.PullRequest) (*models.PullRequest, error)
	GetByReviewer(userID string) ([]models.PullRequest, error)
}
