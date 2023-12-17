package main

import "fmt"

func main() {
	fmt.Println("Welcome to control flows")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "less than 10"
	} else if loginCount > 10 {
		result = "greater than 10"
	} else {
		result = "Equal to 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

	// if err != nil {

	// }
}
