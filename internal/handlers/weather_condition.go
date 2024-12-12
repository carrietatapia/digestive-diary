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

// Example for creating and getting weather conditions
func CreateWeatherConditionHandler(w http.ResponseWriter, r *http.Request) {
	var weatherCondition models.WeatherCondition
	err := json.NewDecoder(r.Body).Decode(&weatherCondition)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO weather_conditions (id, timestamp, temperature, weather_condition, humidity, wind_speed, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		weatherCondition.Timestamp,
		weatherCondition.Temperature,
		weatherCondition.WeatherCondition,
		weatherCondition.Humidity,
		weatherCondition.WindSpeed,
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

func GetWeatherConditionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing weather condition ID", http.StatusBadRequest)
		return
	}

	var weatherCondition models.WeatherCondition
	query := `SELECT id, timestamp, temperature, weather_condition, humidity, wind_speed, created_at FROM weather_conditions WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&weatherCondition.ID,
		&weatherCondition.Timestamp,
		&weatherCondition.Temperature,
		&weatherCondition.WeatherCondition,
		&weatherCondition.Humidity,
		&weatherCondition.WindSpeed,
		&weatherCondition.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Weather condition not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherCondition)
}
