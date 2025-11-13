package repository

import "avitoTestTask/internal/core/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetByID(userID string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}
