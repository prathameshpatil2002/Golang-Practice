package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	//map[key Type]value Type
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["TS"] = "TypeScript"
	languages["RB"] = "Ruby"

	fmt.Println(languages)

	fmt.Println(languages["JS"])

	delete(languages, "RB")

	fmt.Println(languages)

	//Looping through maps

	for key, value := range languages {
		fmt.Printf("key : %v value : %v\n", key, value)
	}

}
