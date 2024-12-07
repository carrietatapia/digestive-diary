package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"prueba/internal/db"
	"prueba/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Example for creating and getting users
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO users (id, email, password_hash, created_at)
        VALUES ($1, $2, $3, $4) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		user.Email,
		user.PasswordHash,
		now,
	).Scan(&id)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, email, password_hash, created_at FROM users WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
