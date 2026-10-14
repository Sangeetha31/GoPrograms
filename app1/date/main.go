package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()

	fmt.Println(dt)
	fmt.Println(dt.Year())
	fmt.Println(dt.Month())
	fmt.Println(dt.Day())
	fmt.Println(dt.Weekday())
	fmt.Println(dt.Hour())
	fmt.Println(dt.Minute())
	fmt.Println(dt.Second())

	fmt.Println(dt.Add(10250 * time.Hour))
	fmt.Println(dt.Add(-10250 * time.Hour))
	fmt.Println(dt.AddDate(0, 526, 0))
	fmt.Println(dt.AddDate(12, 0, 0))

	fmt.Println(dt)
	fmt.Println(dt.Format("31-12-2026"))
	fmt.Println(dt.Format("02/01/2006"))
	fmt.Println(dt.Format("01/02/2006"))
	fmt.Println(dt.Format("2006-02-01"))
	fmt.Println(dt.Format("02/01/2006 15:04:05"))
	fmt.Println(dt.Format("02/01/2026 15:04:05 PM"))

	dt1, err := time.Parse("02-01-2006", "31-12-2005")
	if err != nil {
		fmt.Println(err)
		return
	}

	diff := dt.Sub(dt1)
	fmt.Println("Date-1:", dt1.Format("02-01-2006 03:04 PM"))
	fmt.Println("Date-2:", dt.Format("02-01-2006 03:04 PM"))
	fmt.Println("Diff Hours:", diff.Hours())
	fmt.Println("Diff Days:", diff.Hours()/24)

}
