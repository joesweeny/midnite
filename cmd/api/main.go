package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joesweeny/midnite/internal/app"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var event app.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	report := app.Report{
		Alert:      event.Amount != "0",
		AlertCodes: []int{100, 200},
		UserID:     event.UserID,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(report); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/event", handler)

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
