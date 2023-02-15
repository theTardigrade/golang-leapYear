package leapYear

import (
	"testing"
	"time"
)

func BenchmarkIs(b *testing.B) {
	currentYear := time.Now().UTC().Year()

	for n := 0; n < b.N; n++ {
		Is(currentYear)
	}
}

func TestIs(t *testing.T) {
	leapYears := []int{
		1960,
		2000,
		1804,
		0,
		4,
		8,
	}

	for _, year := range leapYears {
		if !Is(year) {
			t.Fatalf("expected %d to be a leap year", year)
		}
	}

	commonYears := []int{
		1961,
		1800,
		2002,
		1,
		3,
		5,
	}

	for _, year := range commonYears {
		if Is(year) {
			t.Fatalf("expected %d to be a common year", year)
		}
	}
}
