package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case")

	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) + 1

	fmt.Println("value of dice ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Value is 1. You can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot. Roll dice again")
	default:
		fmt.Println("What was that")
	}
}
