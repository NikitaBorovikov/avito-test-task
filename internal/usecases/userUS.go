package usecases

import (
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/core/repository"
)

type UserUC struct {
	UserRepo repository.UserRepo
}

func NewUserUC(userRepo repository.UserRepo) *UserUC {
	return &UserUC{
		UserRepo: userRepo,
	}
}

func (uc *UserUC) SetUserActive(userID string, isActive bool) (*models.User, error) {
	return uc.UserRepo.SetUserActive(userID, isActive)
}
