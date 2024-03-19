package bug

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

type Handler struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewHandler(config util.Config, store db.Store, maker token.Maker) *Handler {
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
		r.Get("/{bug-id}", handler.get)
		r.Put("/", handler.update)
		r.Delete("/{bug-id}", handler.delete)
		r.Get("/by-profile", handler.getByProfile)
		r.Get("/organisations", handler.getOrganisations)
		r.Get("/organisation/{organisation-id}/teams", handler.getOrganisationTeams)
		r.Get("/organisation/{organisation-id}/teams-by-profile", handler.getOrganisationTeamsByProfile)
	})

	return router
}

func (handler *Handler) checkBugPermissions(ctx context.Context, profile, assigneeTeam, assignedTeam uuid.UUID) error {
	getTeamMember := func(ctx context.Context, args db.GetTeamMemberParams) error {
		_, err := handler.store.GetTeamMember(ctx, args)
		return err
	}

	if getTeamMember(ctx, db.GetTeamMemberParams{
		Team:    assignedTeam,
		Profile: profile,
	}) == nil {
		return nil
	}

	if getTeamMember(ctx, db.GetTeamMemberParams{
		Team:    assigneeTeam,
		Profile: profile,
	}) == nil {
		return nil
	}

	return util.ErrUnauthorised
}
