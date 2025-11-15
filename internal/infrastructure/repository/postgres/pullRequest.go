package postgres

import (
	"avitoTestTask/internal/core/models"

	"gorm.io/gorm"
)

type PullRequestRepo struct {
	db *gorm.DB
}

func NewPullRequestRepo(db *gorm.DB) *PullRequestRepo {
	return &PullRequestRepo{
		db: db,
	}
}

func (r *PullRequestRepo) Create(pr *models.PullRequest) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) GetByReviewer(userID string) ([]models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) GetByID(prID string) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) Merge(prID string) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) Reassign(prID, oldReviewerID, newReviewerID string) (*models.PullRequest, error) {
	return nil, nil
}
