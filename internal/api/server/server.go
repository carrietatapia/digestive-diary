package server

import (
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	Router chi.Router // Using the interface here
	DB     *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *Server {
	return &Server{
		DB:     db,
		Router: chi.NewRouter(), // Assign the implementation
	}
}

func (s *Server) Start(port string) error {
	log.Printf("Server started at %s", port)
	return http.ListenAndServe(port, s.Router)
}
