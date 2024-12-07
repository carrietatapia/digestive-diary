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

func CreateEntityHandler(w http.ResponseWriter, r *http.Request) {
	var entity models.Entity
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO entities (id, user_id, name, type, species_id, birth_date, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		entity.UserID,
		entity.Name,
		entity.Type,
		entity.SpeciesID,
		entity.BirthDate,
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

func GetEntityHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing entity ID", http.StatusBadRequest)
		return
	}

	log.Println("entity id:", id)

	var entity models.Entity
	query := `SELECT id, user_id, name, type, species_id, birth_date, created_at FROM entities WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&entity.ID,
		&entity.UserID,
		&entity.Name,
		&entity.Type,
		&entity.SpeciesID,
		&entity.BirthDate,
		&entity.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}
