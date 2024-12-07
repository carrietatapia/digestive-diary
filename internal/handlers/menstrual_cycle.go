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

func CreateMenstrualCycleHandler(w http.ResponseWriter, r *http.Request) {
	var cycle models.MenstrualCycle
	err := json.NewDecoder(r.Body).Decode(&cycle)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	query := `
        INSERT INTO menstrual_cycles (id, entity_id, cycle_start_date, phase, symptoms, created_at)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
    `
	id := uuid.New()
	now := time.Now()
	err = db.DB.QueryRow(
		context.Background(),
		query,
		id,
		cycle.EntityID,
		cycle.CycleStartDate,
		cycle.Phase,
		cycle.Symptoms,
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

func GetMenstrualCycleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing menstrual cycle ID", http.StatusBadRequest)
		return
	}

	var cycle models.MenstrualCycle
	query := `SELECT id, entity_id, cycle_start_date, phase, symptoms, created_at FROM menstrual_cycles WHERE id=$1`
	err := db.DB.QueryRow(context.Background(), query, id).Scan(
		&cycle.ID,
		&cycle.EntityID,
		&cycle.CycleStartDate,
		&cycle.Phase,
		&cycle.Symptoms,
		&cycle.CreatedAt,
	)
	if err != nil {
		log.Println("error :", err)
		http.Error(w, "Menstrual cycle not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cycle)
}
