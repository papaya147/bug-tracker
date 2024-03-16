package team

import (
	"github.com/go-chi/chi/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/util"
)

type Handler struct {
	config util.Config
	store  db.Store
}

func NewHandler(config util.Config, store db.Store) *Handler {
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
