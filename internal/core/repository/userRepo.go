package repository

import "avitoTestTask/internal/core/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	GetByID(userID string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	SetUserActive(userID string, isActive bool) (*models.User, error)
}
