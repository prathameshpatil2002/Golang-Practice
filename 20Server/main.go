package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformGetRequest()
	PerformPostJsonReq()
	PerformPostFormReq()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:1111/get"

	res, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Staus code : ", res.StatusCode)
	fmt.Println("Content len : ", res.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(res.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount : ", byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(string(content))

}

func PerformPostJsonReq() {
	const myUrl = "http://localhost:1111/post"

	//fake payload

	requestBody := strings.NewReader(`
		{
			"courseName":"fake course",
			"price":0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormReq() {
	const myUrl = "http://localhost:1111/postform"

	//formData

	data := url.Values{}
	data.Add("firstName", "Prathm")
	data.Add("lastName", "Patil")
	data.Add("email", "pp@emial.com")

	res, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}
