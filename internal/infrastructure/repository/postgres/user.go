package postgres

import "avitoTestTask/internal/core/models"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return nil
}

func (r *UserRepo) GetByID(userID string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepo) Update(user *models.User) (*models.User, error) {
	return nil, nil
}
