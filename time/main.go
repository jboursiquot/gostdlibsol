package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	aboutTime()
	// duration()
	// epoch()
	// parseAndFormat()
	// timers()
	// afterFunc()
	// timezones()
}

func aboutTime() {
	fmt.Printf("now: %v\n", time.Now())

	t := time.Date(2017, time.August, 1, 18, 25, 45, 0, time.UTC)
	fmt.Printf("past: %v\n", t)
	fmt.Printf("weekday: %s\n", t.Weekday())
	fmt.Printf("month: %d\n", t.Month())
	fmt.Printf("day: %d\n", t.Day())
	fmt.Printf("year: %d\n", t.Year())
	fmt.Printf("hour: %d\n", t.Hour())
	fmt.Printf("minute: %d\n", t.Minute())
	fmt.Printf("second: %d\n", t.Second())
	fmt.Printf("nanosecond: %d\n", t.Nanosecond())
	fmt.Printf("location: %s\n", t.Location())
}

func duration() {
	t1 := time.Now()
	fmt.Printf("t1: %s\n", t1)

	d := 1 * time.Hour
	t2 := t1.Add(d)
	fmt.Printf("t2: %s\n", t2)
	fmt.Printf("t2.Before(t1): %t\n", t2.Before(t1))
	fmt.Printf("t2.After(t1): %t\n", t2.After(t1))
	fmt.Printf("t2.Equal(t1): %t\n", t2.Equal(t1))

	d, _ = time.ParseDuration("60m")
	t3 := t2.Add(-d)
	fmt.Printf("t3: %s\n", t3)
	fmt.Printf("t3.Before(t1): %t\n", t3.Before(t1))
	fmt.Printf("t3.After(t1): %t\n", t3.After(t1))
	fmt.Printf("t3.Equal(t1): %t\n", t3.Equal(t1))
}

func epoch() {
	t := time.Now()
	fmt.Printf("Seconds since Jan 1, 1970 UTC: %d\n", t.Unix())
	fmt.Printf("Nanoseconds since Jan 1, 1970 UTC: %d\n", t.UnixNano())
}

func parseAndFormat() {
	s := "2017-08-01T10:15:55+00:00"
	t, _ := time.Parse(time.RFC3339, s)

	fmt.Println(t)
	fmt.Printf("weekday: %s\n", t.Weekday())
	fmt.Printf("month: %d\n", t.Month())
	fmt.Printf("day: %d\n", t.Day())
	fmt.Printf("year: %d\n", t.Year())
	fmt.Printf("hour: %d\n", t.Hour())
	fmt.Printf("minute: %d\n", t.Minute())
	fmt.Printf("second: %d\n", t.Second())
	fmt.Printf("nanosecond: %d\n", t.Nanosecond())
	fmt.Printf("location: %s\n", t.Location())

	fmt.Printf("RFC3339: %s\n", t.Format(time.RFC3339))
	fmt.Printf("ANSIC: %s\n", t.Format(time.ANSIC))
	fmt.Printf("RubyDate: %s\n", t.Format(time.RubyDate))

	fmt.Printf("Custom: %s\n", t.Format("Jan _2, 2006 (Mon) @ 15:04:05 MST"))
}

func timers() {
	t1 := time.Now()
	tmr := time.NewTimer(3 * time.Second)
	t2 := <-tmr.C // this blocks until 3 seconds have elapsed
	fmt.Printf("%v seconds have gone by\n", t2.Sub(t1).Seconds())
}

func afterFunc() {
	c := make(chan time.Time)
	fmt.Println(time.Now())
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("I'm done")
		c <- time.Now()
	})
	fmt.Println(<-c)
}

func timezones() {
	zones := []string{
		"UTC",
		"Africa/Johannesburg",
		"Africa/Lusaka",
		"America/Port-au-Prince",
		"America/Santiago",
		"America/Toronto",
		"Asia/Seoul",
		"Australia/Melbourne",
		"Europe/Dublin",
		"Europe/Warsaw",
	}

	t := time.Now()
	fmt.Printf("%s - %s\n", t, t.Location())

	for _, tz := range zones {
		l, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s - %s\n", t.In(l), tz)
	}
}
