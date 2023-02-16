# golang-leapYear

This is a simple Go package that deals with leap years.

The `Is` function determines whether a given year contains a leap day or not.

There are also `Next` and `Prev` functions: the former attempts to find the next leap year after a given year, and the latter attempts to find the previous leap year before a given year.

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-leapYear.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-leapYear) [![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-leapYear)](https://goreportcard.com/report/github.com/theTardigrade/golang-leapYear)

## What Is a Leap Year?

In the Gregorian calendar, a leap year has 366 days, whereas other years have only 365. The extra day is known as a leap day.

Approximately one out of every four years is a leap year. Every other year is known as a common year.

For more information, please read [this blog post](https://golangprojectstructure.com/work-out-if-a-year-is-a-leap-year/).

## Example

```golang
package main

import (
	"time"
	"fmt"

	leapYear "github.com/theTardigrade/golang-leapYear"
)

func main() {
	fmt.Println(leapYear.Is(1960)) // true
	fmt.Println(leapYear.Is(1961)) // false
	fmt.Println(leapYear.Is(2000)) // true
	fmt.Println(leapYear.Is(1800)) // false
	fmt.Println(leapYear.Is(1804)) // true

	fmt.Println(leapYear.Next(2020)) // 2024, true
	fmt.Println(leapYear.Next(1977)) // 1980, true

	fmt.Println(leapYear.Prev(2020)) // 2016, true
	fmt.Println(leapYear.Prev(1977)) // 1976, true
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)