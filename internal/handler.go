// package internal

// import (
// 	"encoding/json"
// 	"net/http"
// )

// var counterService *CounterService

// type CounterHandler interface {
// 	IncrementHandler(w http.ResponseWriter, r *http.Request)
// 	DecrementHandler(w http.ResponseWriter, r *http.Request)
// }

// type CounterResponse struct {
// 	Counter int `json:"counter"` // Counter alanı JSON'da "counter" anahtarına sahip olacak anlamına geliyor.
// }

// func IncrementHandler(w http.ResponseWriter, r *http.Request) {
// 	counterService.Increment()
// 	sendCounterResponse(w)
// }

// func DecrementHandler(w http.ResponseWriter, r *http.Request) {
// 	counterService.Decrement()
// 	sendCounterResponse(w)
// }

// func sendCounterResponse(w http.ResponseWriter) {
// 	response := CounterResponse{Counter: counterService.GetCounter()}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// // "w http.ResponseWriter" ifadesi, HTTP yanıtını oluşturmak ve yanıtı istemciye göndermek için kullanılan bir arayüzdür.
// // "r *http.Request" ifadesi, HTTP isteğini temsil eder.

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
