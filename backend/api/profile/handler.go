package profile

import (
	"github.com/go-chi/chi/v5"
	db "github.com/papaya147/bug-tracker/backend/db/sqlc"
)

type Handler struct {
	store *db.Store
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Post("/", handler.create)

	return router
}
