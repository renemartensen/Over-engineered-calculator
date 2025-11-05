package api

import (
	"github.com/gorilla/mux"
)

// NewRouter sets up the HTTP routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/calculate", CalculateHandler).Methods("POST")
	r.HandleFunc("/history", HistoryHandler).Methods("GET") // optional history endpoint
	return r
}
