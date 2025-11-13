package postgres

import "avitoTestTask/internal/core/models"

type PullRequestRepo struct {
}

func NewPullRequestRepo() *PullRequestRepo {
	return &PullRequestRepo{}
}

func (r *PullRequestRepo) Create(pr *models.PullRequest) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) Update(pr *models.PullRequest) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) GetByReviewer(userID string) ([]models.PullRequest, error) {
	return nil, nil
}