package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

const dateFormat = "2006-01-02"

type NewDate struct {
	time.Time
}

func (t NewDate) sundayWeek() int {
	y, _, _ := t.Date()
	beginningOfYear := time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())

	yearDay := float64(t.YearDay())
	weekDayOfFirstDay := float64(beginningOfYear.Weekday())
	ic := yearDay - 7 + weekDayOfFirstDay

	return int(math.Ceil(ic/7) + 1)
}

func calculateRange(start time.Time, end time.Time) []NewDate {
	var result []NewDate

	for d := start; d != end; d = d.AddDate(0, 0, 1) {
		result = append(result, NewDate{d})
	}

	return result
}

func (t NewDate) MarshalJSON() ([]byte, error) {
	type Alias NewDate
	isoYear, isoWeek := t.ISOWeek()
	sundayWeek := t.sundayWeek()

	return json.Marshal(&struct {
		Date           string `json:"date"`
		Year           int    `json:"year"`
		Month          int    `json:"month"`
		Week           int    `json:"week"`
		YearWeek       string `json:"yearWeek"`
		Day            int    `json:"day"`
		IsoYear        int    `json:"isoYear"`
		IsoWeek        int    `json:"isoWeek"`
		IsoYearIsoWeek string `json:"isoYearIsoWeek"`
		Weekday        int    `json:"weekday"`
		Yearday        int    `json:"yearday"`
	}{
		Date:           t.Format(dateFormat),
		Year:           int(t.Year()),
		Month:          int(t.Month()),
		Week:           sundayWeek,
		YearWeek:       fmt.Sprintf("%d%02d", t.Year(), sundayWeek),
		IsoWeek:        isoWeek,
		Day:            int(t.Day()),
		IsoYear:        isoYear,
		IsoYearIsoWeek: fmt.Sprintf("%d%02d", isoYear, isoWeek),
		Weekday:        int(t.Weekday()),
		Yearday:        int(t.YearDay()),
	})
}

func handler(w http.ResponseWriter, r *http.Request) {

	startDate := r.URL.Query().Get("startDate")
	startTime, err := time.Parse(dateFormat, startDate)
	if err != nil {
		http.Error(w, "You need to send a valid startDate.", 400)
		return
	}

	endDate := r.URL.Query().Get("endDate")
	endTime, err := time.Parse(dateFormat, endDate)
	if err != nil {
		http.Error(w, "You need to send a valid endDate.", 400)
		return
	}

	result := calculateRange(startTime, endTime)

	resulArray, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "JSON Error", 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resulArray)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
