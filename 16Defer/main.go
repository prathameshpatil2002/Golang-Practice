package main

import "fmt"

func main() {
	defer fmt.Println("First line") //defer used - will be execute at last - defer execute in last in first out manner
	fmt.Println("Second line")
	defer fmt.Println("Third line")
	fmt.Println("last line")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
