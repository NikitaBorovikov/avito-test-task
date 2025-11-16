package handlers

import (
	apperrors "avitoTestTask/internal/appErrors"
	"avitoTestTask/internal/infrastructure/transport/http/dto"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handlers) GetReviewerStats(w http.ResponseWriter, r *http.Request) {
	reviewerStats, err := h.StatsUC.GetReviewerStats()
	if err != nil {
		logrus.Errorf("failed to get reviewer stats: %v", err)
		sendErrorResponse(w, r, http.StatusInternalServerError, apperrors.ErrorCodeInternalError, apperrors.ErrInternalMsg)
		return
	}

	resp := dto.NewGetReviewerStatsResponse(reviewerStats)
	sendOkResponse(w, r, http.StatusOK, resp)
}
