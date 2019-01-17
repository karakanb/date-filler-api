package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func calculateUniqueRange(start time.Time, end time.Time, uniqueKey string) []Date {
	dates := calculateRange(start, end)

	if uniqueKey != "" {
		dates = uniqueDatesByKey(dates, uniqueKey)
	}

	return dates
}

func handler(w http.ResponseWriter, r *http.Request) {
	startDate, endDate, unique, err := validate(r)
	if err != nil {
		jsonError(w, "You need to send a valid startDate.", 400)
		return
	}

	result := calculateUniqueRange(startDate, endDate, unique)
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting the server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
