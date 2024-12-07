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

// Example for creating and getting meals
func CreateMealHandler(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal
	err := json.NewDecoder(r.Body).Decode(&meal)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO meals (id, entity_id, meal_time, meal_type, foods, beverages, medications, symptoms, mood, weather_id, exercise, exercise_duration, sleep_duration, menstrual_phase, notes, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		meal.EntityID,
		meal.MealTime,
		meal.MealType,
		meal.Foods,
		meal.Beverages,
		meal.Medications,
		meal.Symptoms,
		meal.Mood,
		meal.WeatherID,
		meal.Exercise,
		meal.ExerciseDuration,
		meal.SleepDuration,
		meal.MenstrualPhase,
		meal.Notes,
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

func GetMealHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing meal ID", http.StatusBadRequest)
		return
	}

	var meal models.Meal
	query := `SELECT id, entity_id, meal_time, meal_type, foods, beverages, medications, symptoms, mood, weather_id, exercise, exercise_duration, sleep_duration, menstrual_phase, notes, created_at FROM meals WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&meal.ID,
		&meal.EntityID,
		&meal.MealTime,
		&meal.MealType,
		&meal.Foods,
		&meal.Beverages,
		&meal.Medications,
		&meal.Symptoms,
		&meal.Mood,
		&meal.WeatherID,
		&meal.Exercise,
		&meal.ExerciseDuration,
		&meal.SleepDuration,
		&meal.MenstrualPhase,
		&meal.Notes,
		&meal.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Meal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meal)
}
