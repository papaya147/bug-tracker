package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/papaya147/buggy/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *server) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/ping"),
		middleware.Logger,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/profile", app.profileHandler.Routes())
		r.Mount("/organisation", app.organisationHandler.Routes())
		r.Mount("/team-member", app.teamMemberHandler.Routes())
	})

	router.Mount("/swagger", httpSwagger.WrapHandler)

	return router
}
