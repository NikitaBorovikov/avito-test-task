package handlers

import "avitoTestTask/internal/usecases"

type Handlers struct {
	UserUC        *usecases.UserUC
	TeamUC        *usecases.TeamUC
	PullRequestUC *usecases.PullRequestUC
}

func NewHandlers(uc *usecases.UseCases) *Handlers {
	return &Handlers{
		UserUC:        uc.UserUC,
		TeamUC:        uc.TeamUC,
		PullRequestUC: uc.PullRequestUC,
	}
}
