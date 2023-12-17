package main

import "fmt"

//first letter capital = public variable
const LoginToken string = "uosdfbsdj"

func main() {
	var username string = "Prathamesh"

	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	//default values

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	//implicit

	var website = "website"
	fmt.Println(website)
	//not allowed website = 3

	// no var style
	// this not allowed outside method
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

}
