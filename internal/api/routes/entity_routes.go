package routes

import (
	"github.com/carrietatapia/digestivediary/internal/handlers"
	chi "github.com/go-chi/chi/v5"
)

func SetupEntityRoutes(r chi.Router) {
	r.Route("/entities", func(r chi.Router) {

		r.Post("/", handlers.CreateEntityHandler)
		r.Get("/{id}", handlers.GetEntityHandler)
	})
}
