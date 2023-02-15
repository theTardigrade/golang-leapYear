package leapYear

import "math"

func Next(givenYear int) (foundYear int, found bool) {
	for foundYear = givenYear; foundYear < math.MaxInt; {
		foundYear++

		if found = Is(foundYear); found {
			return
		}
	}

	return
}
