package main

import (
	"fmt"
	"time"
)

func stringToTime(timeStr string) time.Time {
	time, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return time
}
