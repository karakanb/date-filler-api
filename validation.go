package main

import (
	"errors"
	"net/http"
	"time"
)

const startKey = "start"
const endKey = "end"
const uniqueKey = "unique"

var uniqueKeys = map[string]bool{
	"date":           true,
	"year":           true,
	"month":          true,
	"week":           true,
	"yearWeek":       true,
	"day":            true,
	"isoYear":        true,
	"isoWeek":        true,
	"isoYearIsoWeek": true,
	"weekday":        true,
	"yearday":        true,
}

func validate(r *http.Request) (time.Time, time.Time, string, error) {
	startDateString := r.URL.Query().Get(startKey)
	startDate, err := time.Parse(dateFormat, startDateString)
	if err != nil {
		return time.Time{}, time.Time{}, "", errors.New("invalid startDate")
	}

	endDateString := r.URL.Query().Get(endKey)
	endDate, err := time.Parse(dateFormat, endDateString)
	if err != nil {
		return time.Time{}, time.Time{}, "", errors.New("invalid endDate")
	}

	if !endDate.After(startDate) {
		return time.Time{}, time.Time{}, "", errors.New("endDateString has to be after startDateString")
	}

	uniqueKey := r.URL.Query().Get(uniqueKey)
	if _, ok := uniqueKeys[uniqueKey]; uniqueKey != "" && !ok {
		return time.Time{}, time.Time{}, "", errors.New("uniqueKey must be one of the fields in json")
	}

	return startDate, endDate, uniqueKey, nil
}
