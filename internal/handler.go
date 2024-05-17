package internal

import (
	"encoding/json"
	"net/http"
)

type CounterHandler interface {
	IncrementHandler(w http.ResponseWriter, r *http.Request)
	DecrementHandler(w http.ResponseWriter, r *http.Request)
}

type CounterResponse struct {
	Counter int `json:"counter"`
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	CounterServiceInstance.Increment()
	sendCounterResponse(w)
}

func DecrementHandler(w http.ResponseWriter, r *http.Request) {
	CounterServiceInstance.Decrement()
	sendCounterResponse(w)
}

func sendCounterResponse(w http.ResponseWriter) {
	response := CounterResponse{Counter: CounterServiceInstance.GetCounter()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
