package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    // Expected format "7/13/2020 20:32:00"
    format := "1/2/2006 15:04:05"
    t, err := time.Parse(format, date)

    if err != nil {
        fmt.Println("Error parsing date in Schedule:", err)
    }

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    format := "January 2, 2006 15:04:05"
    t, err := time.Parse(format, date)
    if err != nil {
        fmt.Println("Error parsing date in HasPassed:", err)
    }
    return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    format := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(format, date)
    if err != nil {
        fmt.Println("Error parsing date in IsAfternoonAppointment:", err)
    }

    return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    format := "1/2/2006 15:04:05"
    t, err := time.Parse(format, date)
    if err != nil {
        fmt.Println("Error parsing date in Description:", err)
    }

    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
