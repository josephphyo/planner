package main

import (
	"fmt"

	"github.com/josephphyo/planner/dates"
)

func main() {
	fmt.Print("Would you like to convert WeekToDay (or) DayToWeek \n")
	fmt.Print("If you want to convert WeekToDay .. type wtd \n")
	fmt.Print("If you want to convert DayToWeek .. type dtw \n")
	for true {
		fmt.Print("Your input wtd(or)dtw = ")

		var input string
		fmt.Scan(&input)

		switch input {
		case "wtd":
			var week int
			fmt.Print("How many days in week: ")
			fmt.Scan(&week)
			fmt.Println(dates.WeekToDays(week))
		case "dtw":
			var day int
			fmt.Print("How many week in days: ")
			fmt.Scan(&day)
			fmt.Println(dates.DayToWeek(day))
		case "quit":
			return
		}
	}
}
