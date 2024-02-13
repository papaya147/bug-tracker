package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/api/profile"
	"github.com/papaya147/buggy/backend/config"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
)

type server struct {
	config         *config.Config
	store          db.Store
	tokenMaker     token.Maker
	profileHandler *profile.Handler
	router         *chi.Mux
}

func NewServer(store db.Store, maker token.Maker) *server {
	config := config.NewConfig(".")

	server := &server{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}

	server.profileHandler = profile.NewHandler(config, store, maker)
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
