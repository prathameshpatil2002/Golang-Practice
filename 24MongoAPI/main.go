package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prathameshpatil2002/24MongoAPI/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at http://localhost:4000 ...")
}
