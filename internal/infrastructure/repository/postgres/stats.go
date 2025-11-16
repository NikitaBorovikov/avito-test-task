package postgres

import (
	"avitoTestTask/internal/core/models"

	"gorm.io/gorm"
)

type StatsRepo struct {
	db *gorm.DB
}

func NewStatsRepo(db *gorm.DB) *StatsRepo {
	return &StatsRepo{
		db: db,
	}
}

func (r *StatsRepo) GetReviewerStats() ([]models.ReviewerStats, error) {
	var reviewerStats []models.ReviewerStats

	err := r.db.Raw(`
		SELECT u.id as user_id, u.name as user_name, COUNT(pr.id) as count
        FROM users u
        LEFT JOIN pull_requests pr ON u.id = ANY(pr.reviewers)
        WHERE u.is_active = true
        GROUP BY u.id, u.name
        ORDER BY count DESC
	`).Scan(&reviewerStats).Error
	if err != nil {
		return nil, err
	}
	return reviewerStats, nil
}
