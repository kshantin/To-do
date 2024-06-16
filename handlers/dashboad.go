package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"todo/database"
	"todo/models"
)

// func GetNote(http.ResponseWriter, *http.Request)     {}
func GetNotes(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, err := db.Query(context.Background(), `SELECT user_id, name, status, created_at FROM note`)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
	}
	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note
		err := rows.Scan(
			&note.Id,
			&note.UserId,
			&note.Name,
			&note.Status,
			&note.CreatedAt,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Row iteration error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error: %v", err), http.StatusInternalServerError)
	}
}

func AddNote(http.ResponseWriter, *http.Request)     {}
func UpdateNotes(http.ResponseWriter, *http.Request) {}
func DeleteNotes(http.ResponseWriter, *http.Request) {}
