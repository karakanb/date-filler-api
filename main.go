package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func calculateRange(start time.Time, end time.Time) []Date {
	var result []Date

	for d := start; d != end; d = d.AddDate(0, 0, 1) {
		result = append(result, Date{d})
	}

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	startDateString := r.URL.Query().Get("startDate")
	startDate, err := time.Parse(dateFormat, startDateString)
	if err != nil {
		jsonError(w, "You need to send a valid startDate.", 400)
		return
	}

	endDateString := r.URL.Query().Get("endDate")
	endDate, err := time.Parse(dateFormat, endDateString)
	if err != nil {
		jsonError(w, "You need to send a valid endDate.", 400)
		return
	}

	if !endDate.After(startDate) {
		jsonError(w, "endDateString has to be after startDateString", 400)
		return
	}

	result := calculateRange(startDate, endDate)

	resulArray, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "JSON Error", 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resulArray)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting the server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
