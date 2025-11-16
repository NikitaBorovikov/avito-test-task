package postgres

import (
	"avitoTestTask/internal/core/models"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrDublicatePRID = errors.New("pull request with this ID already exists")
	ErrPRNotFound    = errors.New("pull request with this ID is not found")
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
	if err := r.db.Create(pr).Error; err != nil {
		if isDublicateError(err) {
			return nil, ErrDublicatePRID
		}
		return nil, err
	}
	return pr, nil
}

func (r *PullRequestRepo) GetByReviewer(userID string) ([]models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) GetByID(prID string) (*models.PullRequest, error) {
	var pr models.PullRequest
	if err := r.db.Where("id = ?", prID).First(&pr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPRNotFound
		}
		return nil, err
	}
	return &pr, nil
}

func (r *PullRequestRepo) Merge(prID string) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PullRequestRepo) Reassign(prID, oldReviewerID, newReviewerID string) (*models.PullRequest, error) {
	return nil, nil
}
