package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	//slices : don't mention size
	var list = []int{1, 2, 3}

	fmt.Printf("Type of list %T\n", list)

	//append used to add data in slices
	list = append(list, 4)
	fmt.Println(list)

	//slicing like python
	//list = append(list[1:3])
	fmt.Println(list)

	//using make to define slice

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 236
	highScore[2] = 238
	highScore[3] = 240

	fmt.Println(highScore)

	highScore = append(highScore, 658, 585, 694)

	fmt.Println(highScore)

	//sorting
	sort.Ints(highScore)

	fmt.Println(highScore)

	//is sorted or not ?
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slice

	var index int = 2

	highScore = append(highScore[:index], highScore[index+1:]...)

	fmt.Println(highScore)
}
