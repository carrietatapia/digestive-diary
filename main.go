package main

import (
	"log"
	"net/http"

	"github.com/carrietatapia/digestivediary/internal/db"
	"github.com/carrietatapia/digestivediary/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var err error
	db.DB, err = db.ConnectDB("postgres://postgres:admin@localhost:5432/test")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.DB.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/entities", func(r chi.Router) {
		r.Post("/", handlers.CreateEntityHandler)
		r.Get("/{id}", handlers.GetEntityHandler)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.CreateUserHandler)
		r.Get("/{id}", handlers.GetUserHandler)
	})

	r.Route("/species", func(r chi.Router) {
		r.Post("/", handlers.CreateSpeciesHandler)
		r.Get("/{id}", handlers.GetSpeciesHandler)
	})

	r.Route("/stool_consistencies", func(r chi.Router) {
		r.Post("/", handlers.CreateStoolConsistencyHandler)
		r.Get("/{id}", handlers.GetStoolConsistencyHandler)
	})

	r.Route("/weather_conditions", func(r chi.Router) {
		r.Post("/", handlers.CreateWeatherConditionHandler)
		r.Get("/{id}", handlers.GetWeatherConditionHandler)
	})

	r.Route("/meals", func(r chi.Router) {
		r.Post("/", handlers.CreateMealHandler)
		r.Get("/{id}", handlers.GetMealHandler)
	})

	r.Route("/stools", func(r chi.Router) {
		r.Post("/", handlers.CreateStoolHandler)
		r.Get("/{id}", handlers.GetStoolHandler)
	})

	r.Route("/exercises", func(r chi.Router) {
		r.Post("/", handlers.CreateExerciseHandler)
		r.Get("/{id}", handlers.GetExerciseHandler)
	})

	r.Route("/menstrual_cycles", func(r chi.Router) {
		r.Post("/", handlers.CreateMenstrualCycleHandler)
		r.Get("/{id}", handlers.GetMenstrualCycleHandler)
	})

	r.Route("/meal_symptom_links", func(r chi.Router) {
		r.Post("/", handlers.CreateMealSymptomLinkHandler)
		r.Get("/{id}", handlers.GetMealSymptomLinkHandler)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
