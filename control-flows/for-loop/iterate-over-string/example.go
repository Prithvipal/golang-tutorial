package main

import "fmt"

func main() {

	str := "Hello"

	//Get and Print both index and value
	fmt.Println("======================")
	fmt.Println("Printing index and value")
	for index, value := range str {
		fmt.Printf("Index =%d, char =%#U \n", index, value)
	}

	//Get and Print only Index
	fmt.Println("======================")
	fmt.Println("Printing only Indexes")
	for index := range str {
		fmt.Printf("Index =%d \n", index)
	}

	//Get and Print only value. Using _(blank identifiers) for Index
	fmt.Println("======================")
	fmt.Println("Printing only Values")
	for _, value := range str {
		fmt.Printf(" char =%#U \n", value)
	}
}
