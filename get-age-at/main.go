package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	loc , _:= time.LoadLocation("Asia/Tokyo")
	birthDate := time.Date(1981,6,25,0,0,0,0, loc)
	yesterday       := time.Date(2017,6,24,0,0,0,0, loc)
	today       := time.Date(2017,6,25,0,0,0,0, loc)
	fmt.Printf("You are %s years old on yesterday\n", GetAgeAt(birthDate,yesterday))	
	fmt.Printf("You are %s years old on today\n", GetAgeAt(birthDate,today))
}
// https://stackoverflow.com/questions/36530251/golang-time-since-with-months-and-years
func GetAgeAt(birthDate, date time.Time) string {
	if birthDate.Location() != date.Location() {
		date = date.In(birthDate.Location())
	}
	if birthDate.After(date) {
		birthDate, date = date, birthDate

	}
	y1, M1, d1 := birthDate.Date()
	y2, M2, d2 := date.Date()
	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	return strconv.Itoa(year)
}

