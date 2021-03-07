package util

import (
	"fmt"
	"time"
)

// GetWeekByDateString :nodoc:
func GetWeekByDateString(date string) string {
	if t, err := time.Parse(time.RFC3339, date); err == nil {
		year, week := t.Local().ISOWeek()
		return fmt.Sprintf("%v_%v", week, year)
	}
	return ""
}
