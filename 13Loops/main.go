package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Friday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, value := range days {
		fmt.Printf("Index %v day %v\n", index, value)
	}

	rougeValue := 1

	//similar to while loop
	for rougeValue < 10 {

		if rougeValue == 3 {
			goto lco
		}

		if rougeValue == 5 {
			break
		}

		fmt.Println("Value is ", rougeValue)
		rougeValue++
	}

	//go to

lco:
	fmt.Println("Jumped to lco")
}
