package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/carrietatapia/digestivediary/internal/db"
	"github.com/carrietatapia/digestivediary/internal/domain/models"

	chi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Example for creating and getting species
func CreateSpeciesHandler(w http.ResponseWriter, r *http.Request) {
	var species models.Species
	err := json.NewDecoder(r.Body).Decode(&species)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO species (id, name, description, created_at)
        VALUES ($1, $2, $3, $4) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		species.Name,
		species.Description,
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

func GetSpeciesHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing species ID", http.StatusBadRequest)
		return
	}

	var species models.Species
	query := `SELECT id, name, description, created_at FROM species WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&species.ID,
		&species.Name,
		&species.Description,
		&species.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Species not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(species)
}
