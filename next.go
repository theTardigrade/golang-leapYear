package leapYear

import "math"

// Next returns the next leap year after a given year.
// It also returns a boolean value that is set to false if no
// such year can be found.
func Next(givenYear int) (foundYear int, found bool) {
	for foundYear = givenYear; foundYear < math.MaxInt; {
		foundYear++

		if found = Is(foundYear); found {
			return
		}
	}

	return
}
