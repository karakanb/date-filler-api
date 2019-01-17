package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"
)

type Date struct {
	time.Time
	definition DateDefinition
}

type DateDefinition struct {
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
}

func NewDate(t time.Time) Date {
	isoYear, isoWeek := t.ISOWeek()
	sundayWeek := sundayWeek(t)
	definition := DateDefinition{
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
	}

	return Date{t, definition}
}

const dateFormat = "2006-01-02"

// sundayWeek returns the week count based on every week starting on sunday.
func sundayWeek(t time.Time) int {
	y, _, _ := t.Date()
	beginningOfYear := time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())

	yearDay := float64(t.YearDay())
	weekDayOfFirstDayOfYear := float64(beginningOfYear.Weekday())
	ic := yearDay - 7 + weekDayOfFirstDayOfYear

	return int(math.Ceil(ic/7) + 1)
}

func (t Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.definition)
}

func calculateRange(start time.Time, end time.Time) []Date {
	var result []Date

	for d := start; d != end; d = d.AddDate(0, 0, 1) {
		result = append(result, NewDate(d))
	}

	return result
}

func uniqueDatesByKey(dates []Date, key string) []Date {
	u := make([]Date, 0, len(dates))
	m := make(map[string]bool)

	for _, date := range dates {
		identifier := date.definition.getUniqueValueByKey(key)
		if _, ok := m[identifier]; !ok {
			m[identifier] = true
			u = append(u, date)
		}
	}

	return u
}

func (d DateDefinition) getUniqueValueByKey(key string) string {
	builder := strings.Builder{}
	switch key {
	case "date":
		builder.WriteString(string(d.Date))
	case "year":
		builder.WriteString(string(d.Year))
	case "month":
		builder.WriteString(string(d.Year))
		builder.WriteString(string(d.Month))
	case "week":
		builder.WriteString(string(d.Year))
		builder.WriteString(string(d.Week))
	case "yearWeek":
		builder.WriteString(string(d.Year))
		builder.WriteString(string(d.Week))
	case "day":
		builder.WriteString(string(d.Year))
		builder.WriteString(string(d.Month))
		builder.WriteString(string(d.Day))
	case "isoYear":
		builder.WriteString(string(d.IsoYear))
	case "isoWeek":
		builder.WriteString(string(d.IsoYear))
		builder.WriteString(string(d.IsoWeek))
	case "isoYearIsoWeek":
		builder.WriteString(string(d.IsoYearIsoWeek))
	case "weekday":
		builder.WriteString(string(d.Year))
		builder.WriteString(string(d.Week))
		builder.WriteString(string(d.Day))
	case "yearday":
		builder.WriteString(string(d.Year))
		builder.WriteString(string(d.Yearday))
	}

	return builder.String()
}
