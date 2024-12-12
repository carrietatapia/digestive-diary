package routes

import (
	"github.com/carrietatapia/digestivediary/internal/handlers"
	chi "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupStoolRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/stool", func(r chi.Router) {
		r.Post("/", handlers.CreateStoolHandler)
		r.Get("/{id}", handlers.GetStoolHandler)
	})
}
