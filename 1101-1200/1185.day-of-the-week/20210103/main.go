package main

import "fmt"

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day, month, year int) string {
	days := 0
	// 输入年份之前的年份的天数贡献
	days += 365*(year-1971) + (year-1969)/4
	// 输入年份中，输入月份之前的月份的天数贡献
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}
	// 输入月份中的天数贡献
	days += day
	return week[(days+3)%7]
}

func main() {
	fmt.Println(dayOfTheWeek(14, 6, 2011))
}
