package organisation

import (
	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/api/organisation/team"
	"github.com/papaya147/buggy/backend/api/organisation/transfer"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

type Handler struct {
	config          *util.Config
	store           db.Store
	tokenMaker      token.Maker
	transferHandler *transfer.Handler
	teamHandler     *team.Handler
}

func NewHandler(config *util.Config, store db.Store, maker token.Maker) *Handler {
	return &Handler{
		config:          config,
		store:           store,
		tokenMaker:      maker,
		transferHandler: transfer.NewHandler(config, store),
		teamHandler:     team.NewHandler(config, store),
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Group(func(r chi.Router) {
		r.Use(token.Middleware(handler.tokenMaker, handler.store))
		r.Post("/", handler.create)
		r.Get("/", handler.get)
		r.Put("/", handler.update)

		r.Mount("/transfer", handler.transferHandler.Routes())
		r.Mount("/team", handler.teamHandler.Routes())
	})

	return router
}
