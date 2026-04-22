package main

import "fmt"

type sol int

type report struct {
	sol
	location
	temperature
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %vº C\n", report.temperature.high)
}
