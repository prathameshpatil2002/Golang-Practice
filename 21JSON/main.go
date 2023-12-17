package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` //omitempty - don't show if field is empty
}

func main() {
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"JS", 299, "Learncoding.com", "abc123", []string{"Web-dev"}},
		{"MERN", 399, "Learncoding.com", "ab123", []string{"full-stack"}},
		{"ANGULAR", 199, "Learncoding.com", "abc23", nil},
	}

	//package this data as json

	//finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "courseName": "JS",
                "price": 299,
                "website": "Learncoding.com",
                "tags": ["Web-dev"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid json")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) //used to print interface
	} else {
		fmt.Println("Not valid Json")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v ant type %T\n", k, v, v)
	}
}
