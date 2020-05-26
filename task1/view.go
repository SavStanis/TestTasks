package main

import (
	"fmt"
	"time"
)

// Output information about the current lond weekend
func outputCurrLongWeekend(currentLongWeekend LongWeekend, mapOfHolidays map[string]string) {
	longWeekendTimeStart := stringToTime(currentLongWeekend.Start)
	longWeekendTimeFinish := stringToTime(currentLongWeekend.Finish)

	fmt.Printf("There is long weekend of %s now. It will last %d days: %s %d - %s %d\n",
		mapOfHolidays[currentLongWeekend.HolidayDate],
		currentLongWeekend.Duration,
		longWeekendTimeStart.Month().String(),
		longWeekendTimeStart.Day(),
		longWeekendTimeFinish.Month().String(),
		longWeekendTimeFinish.Day())
}

// Output information about the current single holiday
func outputCurrHoliday(mapOfHolidays map[string]string) {
	fmt.Printf("Today is holiday: %s\n", mapOfHolidays[time.Now().Format(timeLayout)])
}

// Output information about the the next long weekend
func outputNextLongWeekend(mapOfHolidays map[string]string, nextLongWeekend LongWeekend) {

	startTime := stringToTime(nextLongWeekend.Start)
	finishTime := stringToTime(nextLongWeekend.Finish)

	fmt.Printf("Today is not a holiday. The next holiday is %s and it will last %d days: %s %d - %s %d\n",
		mapOfHolidays[nextLongWeekend.HolidayDate],
		nextLongWeekend.Duration,
		startTime.Month().String(),
		startTime.Day(),
		finishTime.Month().String(),
		finishTime.Day())
}

// Output information about the the next single holiday
func outputNextHoliday(mapOfHolidays map[string]string, nextHolidayDate string) {

	nextHolidayTime := stringToTime(nextHolidayDate)

	fmt.Printf("Today is not a holiday.The next holiday is %s (%s %d)\n",
		mapOfHolidays[nextHolidayDate],
		nextHolidayTime.Month().String(),
		nextHolidayTime.Day())
}
