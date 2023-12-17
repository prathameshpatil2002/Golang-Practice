package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	//No inheritance in go; No super or parent like concepts

	prathamesh := User{"Prathamesh", "prathamesh@go.dev", true, 21}

	fmt.Println(prathamesh)

	fmt.Printf("Details : %+v\n", prathamesh)
	fmt.Printf("Details : %v\n", prathamesh)
	fmt.Printf("name : %v\n", prathamesh.Name)

	prathamesh.GetStatus()
	prathamesh.NewMail()
}

//First letter capital = Public

type User struct {
	Name  string
	Email string
	Staus bool
	Age   int
}

//method
func (u User) GetStatus() {
	fmt.Println("Is user Active : ", u.Staus)
}

//Does not change actual email as passed copy of user
func (u User) NewMail() {
	u.Email = "ChangedEmail@com"
	fmt.Println("Email of user : ", u.Email)
}
