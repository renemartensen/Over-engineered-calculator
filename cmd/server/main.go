package main

import (
	"github.com/renemartensen/Over-engineered-calculator/internal/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()

	// Serve static files from /web directory
	fs := http.FileServer(http.Dir("./web"))
	router.PathPrefix("/").Handler(fs)

	log.Println("Calculator server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
