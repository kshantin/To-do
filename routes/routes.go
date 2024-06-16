package routes

import (
	"todo/handlers"

	"github.com/gorilla/mux"
)

func GetRoutes(r mux.Router) {
	r.HandleFunc("/api/dashboard", handlers.GetNotes).Methods("GET")
	// r.HandleFunc("/api/dashboard", handlers.GetNote)
	r.HandleFunc("/api/dashboard/delete", handlers.DeleteNotes)
	r.HandleFunc("/api/dashboard/add", handlers.AddNote)
	r.HandleFunc("/api/dashboard/update", handlers.UpdateNotes)
}
