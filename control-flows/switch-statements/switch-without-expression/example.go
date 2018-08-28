package main

import "fmt"

func main() {
	age := readAge()
	switch {
	case age < 18:
		fmt.Println("Invalid age. Underage. Your age", age)
	case age >= 18 && age <= 70:
		fmt.Println("Valid age. Your age", age)
	default:
		fmt.Println("Invalid age. exceed age. Your age", age)
	}
}

func readAge() int {
	var age int
	fmt.Print("Enter your age : ")
	fmt.Scanf("%d", &age)
	return age
}
