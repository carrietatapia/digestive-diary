package routes

import (
	"github.com/carrietatapia/digestivediary/internal/handlers"
	chi "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupMealRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/meal", func(r chi.Router) {
		r.Post("/", handlers.CreateMealHandler)
		r.Get("/{id}", handlers.GetMealHandler)
	})
}
