package routes

import (
	"github.com/carrietatapia/digestivediary/internal/handlers"
	chi "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupStoolConsistencyRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/stool-consistency", func(r chi.Router) {
		r.Post("/", handlers.CreateStoolConsistencyHandler)
		r.Get("/{id}", handlers.GetStoolConsistencyHandler)
	})
}
