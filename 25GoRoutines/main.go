package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually pointer

func main() {
	//goroutine is created by adding keyword go
	// go greeter("Hello") //we are not waiting to return this thread
	// greeter("World")

	websiteList := []string{
		"htps://lco.dev",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) //add waiting - 1 means one group is out there. can ne more
	}
	wg.Wait() //always at end of method
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() //signal done
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("OOPS in %s\n", endpoint)

	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
