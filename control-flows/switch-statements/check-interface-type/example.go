package main

import "fmt"

func main() {
	var day interface{}
	day = "10.4"
	switch day.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	}
}
