package main

import (
	"golang/internal" 
	"net/http"
)

func main() {
	internal.CounterServiceInstance = internal.NewCounterService()

	http.HandleFunc("/increment", internal.IncrementHandler)
	http.HandleFunc("/decrement", internal.DecrementHandler)
	
	http.ListenAndServe(":8080", nil)
}
