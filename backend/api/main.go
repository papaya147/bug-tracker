package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/api/organisation"
	"github.com/papaya147/buggy/backend/api/profile"
	"github.com/papaya147/buggy/backend/api/teammember"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

type server struct {
	config              *util.Config
	store               db.Store
	tokenMaker          token.Maker
	profileHandler      *profile.Handler
	organisationHandler *organisation.Handler
	teamMemberHandler   *teammember.Handler
	router              *chi.Mux
}

func NewServer(store db.Store, maker token.Maker) *server {
	config := util.NewConfig(".")

	server := &server{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}

	server.profileHandler = profile.NewHandler(config, store, maker)
	server.organisationHandler = organisation.NewHandler(config, store, maker)
	server.teamMemberHandler = teammember.NewHandler(config, store, maker)
	server.router = server.routes()

	return server
}

func (app *server) Start(port int) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app.router,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
