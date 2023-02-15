package leapYear

import "math"

// Prev returns the previous leap year before a given leap year.
// It also returns a boolean value that is set to false if no
// such year can be found.
func Prev(givenYear int) (foundYear int, found bool) {
	for foundYear = givenYear; foundYear > math.MinInt; {
		foundYear--

		if found = Is(foundYear); found {
			return
		}
	}

	return
}
