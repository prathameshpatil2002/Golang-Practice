package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("value of pointer ",ptr)

	myNum := 23

	var ptr = &myNum

	fmt.Println("value of pointer ", ptr)
	fmt.Println("value of pointer ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("value of pointer ", myNum)

}
