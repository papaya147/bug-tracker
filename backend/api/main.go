package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/bug-tracker/backend/api/profile"
	"github.com/papaya147/bug-tracker/backend/config"
	db "github.com/papaya147/bug-tracker/backend/db/sqlc"
)

type Server struct {
	config         *config.Config
	store          *db.Store
	profileHandler *profile.Handler
	router         *chi.Mux
}

func NewServer(store *db.Store) *Server {
	config := config.NewConfig()

	server := &Server{
		config: config,
		store:  store,
	}

	server.profileHandler = profile.NewHandler(store)
	server.router = server.routes()

	return server
}

func (app *Server) Start(port int) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app.router,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
