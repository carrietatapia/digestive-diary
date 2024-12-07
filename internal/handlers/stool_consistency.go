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

// Example for creating and getting stool consistencies
func CreateStoolConsistencyHandler(w http.ResponseWriter, r *http.Request) {
	var stoolConsistency models.StoolConsistency
	err := json.NewDecoder(r.Body).Decode(&stoolConsistency)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO stool_consistencies (id, type_number, description, created_at)
        VALUES ($1, $2, $3, $4) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		stoolConsistency.TypeNumber,
		stoolConsistency.Description,
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

func GetStoolConsistencyHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing stool consistency ID", http.StatusBadRequest)
		return
	}

	var stoolConsistency models.StoolConsistency
	query := `SELECT id, type_number, description, created_at FROM stool_consistencies WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&stoolConsistency.ID,
		&stoolConsistency.TypeNumber,
		&stoolConsistency.Description,
		&stoolConsistency.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Stool consistency not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stoolConsistency)
}
