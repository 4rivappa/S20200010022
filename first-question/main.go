package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/4rivappa/S20200010022/first-question/router"
)

func main() {
	fmt.Println("Hello, there .. 1st question!")

	r := router.Router()

	err := http.ListenAndServe(":4068", r)
	if err != nil {
		log.Fatal(err)
	}
}
