package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang/internal"
)

func TestIncrementHandler(t *testing.T) {

	counterService := internal.NewCounterService()

	initialCounter := counterService.GetCounter()

	req := httptest.NewRequest("GET", "http://localhost:8080/increment", nil)

	w := httptest.NewRecorder()

	internal.IncrementHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response internal.CounterResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	expectedCounter := initialCounter + 1
	if response.Counter != expectedCounter {
		t.Errorf("Expected counter value %d, got %d", expectedCounter, response.Counter)
	}
}

func TestDecrementHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080/decrement", nil)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Call the handler function with the mock request and response
	internal.DecrementHandler(w, req)

	// Check the response status code (should be 200 OK)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var response internal.CounterResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Check the counter value in the response
	expectedCounter := -1 // Decrement işlemi sonrası değer 0 olmalı
	if response.Counter != expectedCounter {
		t.Errorf("Expected counter value %d, got %d", expectedCounter, response.Counter)
	}
}
