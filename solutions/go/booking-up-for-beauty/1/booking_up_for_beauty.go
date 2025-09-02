package booking

import ("time"
        "fmt")

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, _ := time.Parse("1/2/2006 15:04:05", date) // time.Time, error
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsed, _ := time.Parse("January 2, 2006 15:04:05", date)
    var passed bool
    switch parsed.Compare(time.Now()) {
    case -1: 
    	passed = true	
    case 0:
    	passed = true
    default: 
    	passed = false
    }
	return passed
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsed, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if (parsed.Hour() >= 12) && (parsed.Hour() < 18) {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
    mytime := Schedule(date)
    formattedDate := mytime.Format("Monday, January 2, 2006,")
    formattedTime := mytime.Format("15:04")
    return fmt.Sprintf("You have an appointment on %s at %s.", formattedDate, formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t, _ := time.Parse("1/2/2006", "9/15/2012")
	return t.AddDate(time.Now().Year()-2012, 0, 0)
}
