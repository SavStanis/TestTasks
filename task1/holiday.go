package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Holiday struct {
	Date      string
	LocalName string
	Name      string
}

// Get data from API
func getDataFromAPI() []Holiday {
	resp, err := http.Get("https://date.nager.at/api/v2/PublicHolidays/" + Year + "/" + CountryIndex)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer resp.Body.Close()

	var w []Holiday
	json.NewDecoder(resp.Body).Decode(&w)

	return w
}

func getAndParseData() map[string]string {
	data := getDataFromAPI()
	mapOfHolidays := make(map[string]string)
	fillMapOfHolidays(data, mapOfHolidays)

	return mapOfHolidays
}

func fillMapOfHolidays(holidays []Holiday, mapOfHolidays map[string]string) {
	for _, holiday := range holidays {
		mapOfHolidays[holiday.Date] = holiday.Name
	}
}

func isHolidayToday(mapOfHolidays map[string]string) bool {
	now := time.Now()

	if mapOfHolidays[now.Format(timeLayout)] != "" {
		return true
	}
	return false
}

func getNextHolidayDate(mapOfHolidays map[string]string) string {
	now := time.Now()
	var nextHoliday string

	for holiday := range mapOfHolidays {
		currHolidayTime := stringToTime(holiday)

		if currHolidayTime.Before(now) {
			continue
		}

		if nextHoliday == "" {
			nextHoliday = holiday
			continue
		}

		nextHolidayTime := stringToTime(nextHoliday)
		if nextHolidayTime.After(currHolidayTime) {
			nextHoliday = holiday
		}
	}

	return nextHoliday
}
