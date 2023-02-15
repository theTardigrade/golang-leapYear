# golang-leapYear

This is a simple Go package that deals with leap years.

In the Gregorian calendar, a leap year has 366 days, whereas other years have only 365. Approximately one out of every four years is a leap year. Every other year is known as a common year.

There is only one public function, which determines whether a given year contains a leap day or not. 

[![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-leapYear)](https://goreportcard.com/report/github.com/theTardigrade/golang-leapYear)

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
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)