package leapYear

// Is returns true only if the given year contains a leap day,
// meaning that the year is a leap year.
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
