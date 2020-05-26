package main

import (
	"time"
)

type LongWeekend struct {
	HolidayDate string
	Start       string
	Finish      string
	Duration    byte
}

func calculateLongWeekends(mapOfHolidays map[string]string) []LongWeekend {
	var longWeekends []LongWeekend

	for holiday := range mapOfHolidays {
		t := stringToTime(holiday)

		if t.Weekday().String() == "Friday" {
			longWeekend := LongWeekend{t.Format(timeLayout), t.Format(timeLayout), t.AddDate(0, 0, 2).Format(timeLayout), 3}
			longWeekends = append(longWeekends, longWeekend)
		}

		if t.Weekday().String() == "Monday" {
			longWeekend := LongWeekend{t.Format(timeLayout), t.AddDate(0, 0, -2).Format(timeLayout), t.Format(timeLayout), 3}
			longWeekends = append(longWeekends, longWeekend)
		}
	}

	return longWeekends
}



func isLongWeekendNow(longWeekends []LongWeekend) (bool, LongWeekend) {
	now := time.Now()

	for _, longWeekend := range longWeekends {
		startTime := stringToTime(longWeekend.Start)
		finishTime := stringToTime(longWeekend.Start)

		if now.After(startTime) && finishTime.After(now) {
			return true, longWeekend
		}
	}

	return false, LongWeekend{}
}



func getNextLongWeekend(longWeekends []LongWeekend) LongWeekend {
	now := time.Now()
	nextLongWeekend := LongWeekend{}

	for _, currLongWeekend := range longWeekends {

		currLongWeekendTime := stringToTime(currLongWeekend.Start)

		if currLongWeekendTime.Before(now) {
			continue
		}

		if nextLongWeekend.Start == "" || nextLongWeekend.Finish == "" || nextLongWeekend.HolidayDate == "" || nextLongWeekend.Duration == 0 {
			nextLongWeekend = currLongWeekend
			continue
		}

		nextLongWeekendTime := stringToTime(nextLongWeekend.Start)

		if nextLongWeekendTime.After(currLongWeekendTime) {
			nextLongWeekend = currLongWeekend
		}
	}

	return nextLongWeekend
}