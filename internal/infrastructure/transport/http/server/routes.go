package server

import (
	"avitoTestTask/internal/infrastructure/transport/http/handlers"

	"github.com/go-chi/chi/v5"
)

func initRoutes(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()
	return r
}
