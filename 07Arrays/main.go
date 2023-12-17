package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var list = [3]int{1, 2, 3}

	fmt.Println(list)

}
