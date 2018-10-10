package main

import "fmt"

func main() {

	var alphabetArr [4]string
	alphabetArr[0] = "Apple"
	alphabetArr[1] = "Ball"
	alphabetArr[2] = "Cat"
	alphabetArr[3] = "Dog"

	//Get and Print both Index and value
	fmt.Println("======================")
	fmt.Println("Printing Index and value")
	for key, value := range alphabetArr {
		fmt.Println("Index =", key, "Value =", value)
	}

	//Get and Print only Index
	fmt.Println("======================")
	fmt.Println("Printing only Indexes")
	for key := range alphabetArr {
		fmt.Println("Index =", key)
	}

	//Get and Print only value. Using _(blank identifiers) for Indexes
	fmt.Println("======================")
	fmt.Println("Printing only Values")
	for _, value := range alphabetArr {
		fmt.Println("Value =", value)
	}
}
