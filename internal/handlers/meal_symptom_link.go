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

func CreateMealSymptomLinkHandler(w http.ResponseWriter, r *http.Request) {
	var link models.MealSymptomLink
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO meal_symptom_links (id, meal_id, symptom, created_at)
        VALUES ($1, $2, $3, $4) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		link.MealID,
		link.Symptom,
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

func GetMealSymptomLinkHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing meal symptom link ID", http.StatusBadRequest)
		return
	}

	var link models.MealSymptomLink
	query := `SELECT id, meal_id, symptom, created_at FROM meal_symptom_links WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&link.ID,
		&link.MealID,
		&link.Symptom,
		&link.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Meal symptom link not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}
