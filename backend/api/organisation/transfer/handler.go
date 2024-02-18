package transfer

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
		config: config,
		store:  store,
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Post("/", handler.create)
	router.Get("/", handler.get)
	router.Delete("/{organisation-transfer-id}", handler.delete)

	return router
}
