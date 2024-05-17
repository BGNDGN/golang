package main

import (
	"golang/internal" // 'yourmodule' yerine modül adınızı kullanın
	"net/http"
)

func main() {
	// counterService'i başlatın
	internal.CounterServiceInstance = internal.NewCounterService()

	http.HandleFunc("/increment", internal.IncrementHandler)
	http.HandleFunc("/decrement", internal.DecrementHandler)

	// Sunucuyu başlatın
	http.ListenAndServe(":8080", nil)
}
