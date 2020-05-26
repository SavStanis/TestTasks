package main

func main() {

	mapOfHolidays := getAndParseData()
	longWeekends := calculateLongWeekends(mapOfHolidays)

	isLongWeekendNow, currentLongWeekend := isLongWeekendNow(longWeekends)

	if isLongWeekendNow {
		outputCurrLongWeekend(currentLongWeekend, mapOfHolidays)
	} else if isHolidayToday(mapOfHolidays) {
		outputCurrHoliday(mapOfHolidays)
	} else {
		nextHolidayDate := getNextHolidayDate(mapOfHolidays)
		nextLongWeekend := getNextLongWeekend(longWeekends)

		nextHolidayTime := stringToTime(nextHolidayDate)
		nextLongWeekendStartTime := stringToTime(nextLongWeekend.Start)

		if nextLongWeekend.Start == nextHolidayDate || nextLongWeekendStartTime.Before(nextHolidayTime) {
			outputNextLongWeekend(mapOfHolidays, nextLongWeekend)
		} else {
			outputNextHoliday(mapOfHolidays, nextHolidayDate)
		}
	}
}
