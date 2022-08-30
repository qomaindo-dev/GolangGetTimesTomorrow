package main

import (
	"fmt"
	"time"

	"github.com/senseyeio/duration"
)

func main() {
	d, _ := duration.ParseISO8601("P1D")

	timeLocation, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(timeLocation)

	timeFormat := "2006-01-02T15:04:05.999Z07:00"
	timeSToday := timeNow.Format(timeFormat)

	timesTomorrow := d.Shift(timeNow)
	timesTomorrowFormatted := timesTomorrow.Format(timeFormat)

	fmt.Println("Today: ", timeSToday)
	fmt.Println("Tomorrow: ", timesTomorrowFormatted)
}
