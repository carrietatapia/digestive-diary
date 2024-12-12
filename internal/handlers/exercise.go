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

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO exercises (id, entity_id, exercise_type, duration, intensity, performed_at, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		exercise.EntityID,
		exercise.ExerciseType,
		exercise.Duration,
		exercise.Intensity,
		exercise.PerformedAt,
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

func GetExerciseHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing exercise ID", http.StatusBadRequest)
		return
	}

	var exercise models.Exercise
	query := `SELECT id, entity_id, exercise_type, duration, intensity, performed_at, created_at FROM exercises WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&exercise.ID,
		&exercise.EntityID,
		&exercise.ExerciseType,
		&exercise.Duration,
		&exercise.Intensity,
		&exercise.PerformedAt,
		&exercise.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercise)
}
