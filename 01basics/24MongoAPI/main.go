package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com.mayur/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to mongodb api")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))

}
