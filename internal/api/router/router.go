package router

import (
	"github.com/carrietatapia/digestivediary/internal/api/middleware"
	"github.com/carrietatapia/digestivediary/internal/api/routes"
	chi "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Setup routes
	routes.SetupEntityRoutes(r)
	routes.SetupUserRoutes(r)
	//routes.SetupSpeciesRoutes(r, db)
	//routes.SetupStoolConsistencyRoutes(r, db)
	//routes.SetupStoolRoutes(r, db)
	//routes.SetupMealRoutes(r, db)
	////routes.SetupExerciseRoutes(r, db)

	return r
}
