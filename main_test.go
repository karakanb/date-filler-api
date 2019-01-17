package main

import (
	"testing"
	"time"
)

func TestCalculateRange(t *testing.T) {
	tables := []struct {
		startDate        string
		endDate          string
		expectedDayCount int
	}{
		{"2013-12-30", "2014-01-10", 11},
	}

	for _, table := range tables {

		t.Run(table.startDate, func(t *testing.T) {
			startDate, _ := time.Parse(dateFormat, table.startDate)
			endDate, _ := time.Parse(dateFormat, table.endDate)

			result := calculateRange(startDate, endDate)
			dayCount := len(result)

			if dayCount != table.expectedDayCount {
				t.Errorf("calculateRange failed, got: %d days, want: %d days", dayCount, table.expectedDayCount)
			}
		})
	}

}
