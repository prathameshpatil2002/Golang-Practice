package main

import "fmt"

func main() {
	fmt.Println("Welcome to function")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proRes := proAdder(2, 5, 6, 7, 8)

	fmt.Println("Result is: ", proRes)

	demoVal, demoMsg := demo(2)

	fmt.Println(demoVal, demoMsg)

}

//func name (parameters) return_type {}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

//without defining number of values
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

//Multiple return values

func demo(val int) (int, string) {
	return val, "Demo message"
}

func greeter() {
	fmt.Println("Hello I am greeter")
}
