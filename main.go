package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

const PortKey = "PORT"
const DefaultPort = "8080"

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
		jsonError(w, err.Error(), 400)
		return
	}

	result := calculateUniqueRange(startDate, endDate, unique)
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

func getPort() string {
	value := os.Getenv(PortKey)
	if len(value) == 0 {
		value = DefaultPort
	}
	return ":" + value
}

func main() {
	http.HandleFunc("/", handler)
	port := getPort()
	log.Println("Starting the server in port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
