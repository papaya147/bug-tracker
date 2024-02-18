package organisation

import (
	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/config"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
)

type Handler struct {
	config     *config.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewHandler(config *config.Config, store db.Store, maker token.Maker) *Handler {
	return &Handler{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Group(func(r chi.Router) {
		r.Use(token.Middleware(handler.tokenMaker, handler.store))
		r.Post("/", handler.create)
		r.Get("/", handler.get)
		r.Put("/", handler.update)
	})

	return router
}
