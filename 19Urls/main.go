package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=js&paymentId=2"

func main() {

	fmt.Println(myUrl)

	//parsing

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of qparams : %T\n", qparams)
	fmt.Println(qparams)

	//dont forgot &
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/test",
		RawQuery: "user=testUser",
	}

	anotherUel := partsOfUrl.String()

	fmt.Println(anotherUel)

}
