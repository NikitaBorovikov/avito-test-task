package handlers

import (
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"net/http"
)

const (
	userIDQueryParam = "user_id"
)

func (h *Handlers) GetUserReviewPR(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get(userIDQueryParam)

	prs, err := h.PullRequestUC.GetByReviewer(userID)
	if err != nil {
		// TODO: handle error
		return
	}

	resp := dto.NewGetUserReviewPRResponse(userID, prs)
	sendOkResponse(w, r, http.StatusOK, resp)
}

func (h *Handlers) SetUserActive(w http.ResponseWriter, r *http.Request) {

}
