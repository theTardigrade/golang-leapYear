package leapYear

// Is returns true only if the given year contains a leap day.
// In the Gregorian calendar, a leap year has 366 days, whereas
// other years have only 365.
// Aproximately one out of every four years is a leap year.
func Is(givenYear int) bool {
	if givenYear%400 == 0 {
		return true
	} else if givenYear%100 == 0 {
		return false
	} else if givenYear%4 == 0 {
		return true
	}

	return false
}
