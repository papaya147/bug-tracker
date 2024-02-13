package profile

import (
	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/config"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
)

type Handler struct {
	store      db.Store
	config     *config.Config
	tokenMaker token.Maker
}

func NewHandler(store db.Store, config *config.Config, maker token.Maker) *Handler {
	return &Handler{
		store:      store,
		config:     config,
		tokenMaker: maker,
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Post("/", handler.create)
	router.Get("/verify", handler.verify)
	router.Post("/login", handler.login)

	router.Group(func(r chi.Router) {
		r.Use(token.Middleware(handler.tokenMaker, handler.store))
		r.Get("/", handler.get)
	})

	return router
}
