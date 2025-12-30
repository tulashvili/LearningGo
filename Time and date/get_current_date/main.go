package main

import (
	"fmt"
	"time"
)

type Day struct {
	Date time.Time
}

type Date struct {
	Day   int
	Month string
	Year  int
}

func DateWithModel() {
	var date = Date{
		Day:   time.Now().Day(),
		Month: time.Now().Month().String(),
		Year:  time.Now().Year(),
	}
	fmt.Println(date.Day, date.Month, date.Year)
}

//	func TodayWithUTC(t time.Time) Day {
//		y, m, d := t.Date()
//		return Day{
//			Date: time.Date(y, m, d, 2, 0, 0, 0, t.Location()),
//		}
//	}
func TodayWithUTC(t time.Time) Day {
	y, m, d := t.Date()
	return Day{
		Date: time.Date(y, m, d, 0, 0, 0, 0, t.Location()),
	}
}

// func ParseDates() {
// 	today := TodayWithUTC(time.Now())
// 	layout := "2006-01-02 15:04:05 +0300 MSK"
// 	t, err := time.Parse(layout, today.Date.String())
// 	if err != nil {
// 		fmt.Println("Error parsing time:", err)
// 		return
// 	}
// 	fmt.Printf("Parsed time: %v, type: %T\n", t, t)
// }

func main() {
	todayWithUTC := TodayWithUTC(time.Now())
	fmt.Println(todayWithUTC)

}
