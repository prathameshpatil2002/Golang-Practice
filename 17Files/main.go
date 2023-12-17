package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "This is file text"

	file, err := os.Create("./sample.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is ", length)

	defer file.Close()

	readFile("./sample.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text Data inside file :\n", string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err) // shut down programme
	}
}
