package api

import (
	"encoding/json"
	"github.com/renemartensen/Over-engineered-calculator/internal/calculator"
	"github.com/renemartensen/Over-engineered-calculator/internal/storage"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// CalculateHandler handles /calculate POST requests
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "invalid JSON"})
		return
	}

	result, err := calculator.EvaluateExpression(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

	// Store the calculation in history
	storage.MemoryStoreInstance.Add(req.Expression, result)
	json.NewEncoder(w).Encode(Response{Result: result})
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	history := storage.MemoryStoreInstance.GetAll()
	json.NewEncoder(w).Encode(history)
}
