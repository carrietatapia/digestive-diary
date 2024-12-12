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

func CreateStoolHandler(w http.ResponseWriter, r *http.Request) {
	var stool models.Stool
	err := json.NewDecoder(r.Body).Decode(&stool)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO stools (id, entity_id, observed_at, consistency_id, color, volume, mood, weather_id, exercise, exercise_duration, sleep_duration, menstrual_phase, notes, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		stool.EntityID,
		stool.ObservedAt,
		stool.ConsistencyID,
		stool.Color,
		stool.Volume,
		stool.Mood,
		stool.WeatherID,
		stool.Exercise,
		stool.ExerciseDuration,
		stool.SleepDuration,
		stool.MenstrualPhase,
		stool.Notes,
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

func GetStoolHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing stool ID", http.StatusBadRequest)
		return
	}

	var stool models.Stool
	query := `SELECT id, entity_id, observed_at, consistency_id, color, volume, mood, weather_id, exercise, exercise_duration, sleep_duration, menstrual_phase, notes, created_at FROM stools WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&stool.ID,
		&stool.EntityID,
		&stool.ObservedAt,
		&stool.ConsistencyID,
		&stool.Color,
		&stool.Volume,
		&stool.Mood,
		&stool.WeatherID,
		&stool.Exercise,
		&stool.ExerciseDuration,
		&stool.SleepDuration,
		&stool.MenstrualPhase,
		&stool.Notes,
		&stool.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Stool not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stool)
}
