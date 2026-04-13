package booking

import "fmt"
import "time"
import "strconv"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, _ := time.Parse("1/2/2006 15:04:05", date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    appointmentTime, _ := time.Parse(layout, date)
    currentDate := time.Now()
	return currentDate.After(appointmentTime)
}

func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    appointmentTime, _ := time.Parse(layout, date)
    hourStr := appointmentTime.Format("15")
    hour, _ := strconv.Atoi(hourStr)
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    appointmentTime, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s.", appointmentTime.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
