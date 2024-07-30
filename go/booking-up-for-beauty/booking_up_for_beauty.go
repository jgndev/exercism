package booking

import (
	"time"
)

// const layout string = "1/2/2006, 15:04:05"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	pt, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}

	return pt
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	pt, err := time.Parse(layout, date)
	if err != nil {
		// This function signature doesn't return an error, otherwise we would
		// likely return false and the error here: false, err.
		return false
	}

	return time.Now().After(pt)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	pt, err := time.Parse(layout, date)
	if err != nil {
		// This function signature doesn't return an error, otherwise we would
		// likely return false and the error here: false, err.
		return false
	}

	return pt.Hour() >= 12 && pt.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	pt, err := time.Parse(layout, date)
	if err != nil {
		// This function signature doesn't return an error, otherwise we would
		// likely return an empty string and the error here: "", err.
		return ""
	}

	desc := pt.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")

	return desc
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
