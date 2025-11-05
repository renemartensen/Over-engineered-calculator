package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/renemartensen/Over-engineered-calculator/internal/storage"
)

func resetStore() {
	storage.MemoryStoreInstance = &storage.MemoryStore{}
}

func TestCalculateHandler_ValidExpression(t *testing.T) {
	resetStore()

	reqBody := `{"expression":"2+3*4"}`
	req := httptest.NewRequest("POST", "/calculate", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	CalculateHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d, body: %s", resp.StatusCode, body)
	}

	var res map[string]float64
	err := json.Unmarshal(body, &res)
	if err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if res["result"] != 14 { // 2 + 3*4 = 14
		t.Fatalf("expected result 14, got %v", res["result"])
	}
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
	resetStore()

	req := httptest.NewRequest("POST", "/calculate", strings.NewReader(`{invalid json}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	CalculateHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid JSON, got %d", resp.StatusCode)
	}
}

func TestCalculateHandler_InvalidExpression(t *testing.T) {
	resetStore()

	req := httptest.NewRequest("POST", "/calculate", strings.NewReader(`{"expression":"2+&3"}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	CalculateHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid expression, got %d", resp.StatusCode)
	}
}

func TestHistoryHandler(t *testing.T) {
	resetStore()

	// Seed some history
	storage.MemoryStoreInstance.Add("1+1", 2)
	storage.MemoryStoreInstance.Add("3+4", 7)

	req := httptest.NewRequest("GET", "/history", nil)
	w := httptest.NewRecorder()

	HistoryHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var history []map[string]interface{}
	if err := json.Unmarshal(body, &history); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(history) != 2 {
		t.Fatalf("expected 2 history items, got %d", len(history))
	}

	if history[0]["expression"] != "1+1" || history[0]["result"].(float64) != 2 {
		t.Fatalf("first history item mismatch")
	}
}

func TestAuthMiddleware(t *testing.T) {
	// Wrap a simple test handler with your AuthMiddleware
	testHandler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized"))
	}))

	t.Run("Authorized request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.SetBasicAuth("user@example.com", "123456") // correct credentials

		rr := httptest.NewRecorder()
		testHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rr.Code)
		}
		if strings.TrimSpace(rr.Body.String()) != "Authorized" {
			t.Errorf("Expected body 'Authorized', got %s", rr.Body.String())
		}
	})

	t.Run("Unauthorized request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.SetBasicAuth("user@example.com", "wrongpassword") // wrong credentials

		rr := httptest.NewRecorder()
		testHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %d", rr.Code)
		}
	})
}
