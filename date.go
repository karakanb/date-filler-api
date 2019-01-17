package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type Date struct {
	time.Time
}

const dateFormat = "2006-01-02"

// sundayWeek returns the week count based on every week starting on sunday.
func (t Date) sundayWeek() int {
	y, _, _ := t.Date()
	beginningOfYear := time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())

	yearDay := float64(t.YearDay())
	weekDayOfFirstDayOfYear := float64(beginningOfYear.Weekday())
	ic := yearDay - 7 + weekDayOfFirstDayOfYear

	return int(math.Ceil(ic/7) + 1)
}

func (t Date) MarshalJSON() ([]byte, error) {
	type Alias Date
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
