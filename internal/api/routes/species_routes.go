package routes

import (
	"github.com/carrietatapia/digestivediary/internal/handlers"
	chi "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupSpeciesRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/species", func(r chi.Router) {
		r.Post("/", handlers.CreateSpeciesHandler)
		r.Get("/{id}", handlers.GetSpeciesHandler)
	})
}
