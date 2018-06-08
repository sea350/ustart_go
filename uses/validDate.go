package uses

import (
	"fmt"
	"regexp"
	"strconv"
)

//LeapDays ...
//Calculates number of days in February, based on the year
func LeapDays(year string) int {
	leap := 29
	notleap := 28
	inputYear, _ := strconv.Atoi(year)
	if inputYear%4 == 0 {
		return leap
	}
	return notleap

}

var nDays = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

//ValidDate ..
//Date Validation
func ValidDate(date string) bool {

	// 	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	if len(date) < 10 {
		fmt.Println("length mismatch: 11 vs", len(date))
		return false
	}
	rxDate := regexp.MustCompile("(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])/((19|20)\\d\\d)")

	// 06/08/2018
	month, errMonth := strconv.Atoi(date[0:2])
	if errMonth != nil {

		fmt.Println("month atoi")
		return false
	}
	days, errDays := strconv.Atoi(date[3:5])
	if errDays != nil {
		fmt.Println("days atoi")
		return false
	}
	// year, errYear := strconv.Atoi(date[0:4])
	// if errYear != nil {
	// 	return false
	// }

	if date[0:2] == "2" && days > LeapDays(date[6:]) {
		fmt.Println("wrong february days")
		return false
	} else if days > nDays[month] {
		fmt.Println("wrong days for month", days)
		return false
	}

	if !rxDate.MatchString(date) {
		fmt.Println("matchstring error")
		return false
	}
	return true
}
