package team

import (
	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/config"
	db "github.com/papaya147/buggy/backend/db/sqlc"
)

type Handler struct {
	config *config.Config
	store  db.Store
}

func NewHandler(config *config.Config, store db.Store) *Handler {
	return &Handler{
		store:  store,
		config: config,
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Post("/", handler.create)
	router.Get("/", handler.get)
	router.Put("/{team-id}", handler.update)

	return router
}
