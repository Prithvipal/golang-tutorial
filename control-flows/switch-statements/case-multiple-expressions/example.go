package main

import "fmt"

func main() {
	day := readDay()
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}

func readDay() int {
	var day int
	fmt.Print("Enter day : ")
	fmt.Scanf("%d", &day)
	return day
}
