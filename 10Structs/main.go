package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	//No inheritance in go; No super or parent like concepts

	prathamesh := User{"Prathamesh", "prathamesh@go.dev", true, 21}

	fmt.Println(prathamesh)

	fmt.Printf("Details : %+v\n", prathamesh)
	fmt.Printf("Details : %v\n", prathamesh)
}

//First letter capital = Public

type User struct {
	Name  string
	Email string
	Staus bool
	Age   int
}
