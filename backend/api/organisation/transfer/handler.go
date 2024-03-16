package transfer

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
		config: config,
		store:  store,
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Post("/", handler.create)
	router.Get("/", handler.get)
	router.Delete("/{organisation-transfer-id}", handler.delete)
	router.Get("/response/{organisation-transfer-id}", handler.response)

	return router
}
