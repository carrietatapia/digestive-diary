package routes

import (
	"github.com/carrietatapia/digestivediary/internal/db"
	"github.com/carrietatapia/digestivediary/internal/handlers"
	"github.com/carrietatapia/digestivediary/internal/repository/postgres"
	"github.com/carrietatapia/digestivediary/internal/service/impl"
	chi "github.com/go-chi/chi/v5"
)

func SetupUserRoutes(r chi.Router) {

	userService := impl.NewUserServiceImpl(postgres.NewUserRepository(db.DB))
	userHandler := handlers.NewUserHandler(userService)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUserHandler)
		r.Get("/{id}", userHandler.GetUserHandler)
		r.Put("/{id}", userHandler.UpdateUserHandler)
		r.Delete("/{id}", userHandler.DeleteUserHandler)
	})
}
