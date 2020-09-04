// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/RadhiFadlillah/go-prayer"
)

func main() {
	// Prepare calculator
	calc := prayer.Calculator{
		Latitude:          -6.21,
		Longitude:         106.85,
		Elevation:         50,
		CalculationMethod: prayer.Kemenag,
		AsrConvention:     prayer.Shafii,
		PreciseToSeconds:  false,
		TimeCorrection: prayer.TimeCorrection{
			prayer.Fajr:    2 * time.Minute,
			prayer.Sunrise: -time.Minute,
			prayer.Zuhr:    2 * time.Minute,
			prayer.Asr:     time.Minute,
			prayer.Maghrib: time.Minute,
			prayer.Isha:    time.Minute,
		},
	}

	// Initiate the calculator. It must be run every time any
	// parameter in calculator above changed.
	calc.Init()

	// Specify the date that you want to calculate. This package
	// will use timezone from your date, so make sure to set it
	// to the timezone of the location that you want to calculate.
	zone := time.FixedZone("WIB", 7*3600)
	date := time.Date(2020, 9, 4, 0, 0, 0, 0, zone)
	calc.SetDate(date)

	// Now we just need to calculate
	result := calc.Calculate()
	fmt.Println(date.Format("2006-01-02"))
	fmt.Println("Fajr    =", result[prayer.Fajr].Format("15:04"))
	fmt.Println("Sunrise =", result[prayer.Sunrise].Format("15:04"))
	fmt.Println("Zuhr    =", result[prayer.Zuhr].Format("15:04"))
	fmt.Println("Asr     =", result[prayer.Asr].Format("15:04"))
	fmt.Println("Maghrib =", result[prayer.Maghrib].Format("15:04"))
	fmt.Println("Isha    =", result[prayer.Isha].Format("15:04"))
}
