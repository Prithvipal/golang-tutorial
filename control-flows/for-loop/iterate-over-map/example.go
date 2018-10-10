package main

import "fmt"

func main() {
	alphabetMap := map[string]string{
		"A": "Apple",
		"B": "Ball",
		"C": "Cat",
		"D": "Dog",
	}

	//Get and Print both key and value
	fmt.Println("======================")
	fmt.Println("Printing key and value")
	for key, value := range alphabetMap {
		fmt.Println("Key =", key, "Value =", value)
	}

	//Get and Print only key
	fmt.Println("======================")
	fmt.Println("Printing only keys")
	for key := range alphabetMap {
		fmt.Println("Key =", key)
	}

	//Get and Print only value. Using _(blank identifiers) for keys
	fmt.Println("======================")
	fmt.Println("Printing only Values")
	for _, value := range alphabetMap {
		fmt.Println("Value =", value)
	}
}
