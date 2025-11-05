package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(AuthMiddleware)

	protected.HandleFunc("/calculate", CalculateHandler).Methods("POST")
	protected.HandleFunc("/history", HistoryHandler).Methods("GET")

	return router
}
