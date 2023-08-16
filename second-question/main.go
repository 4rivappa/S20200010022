package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/4rivappa/S20200010022/second-question/router"
)

func main() {
	fmt.Println("Hello, there !")

	r := router.Router()

	err := http.ListenAndServe(":4068", r)
	if err != nil {
		log.Fatal(err)
	}
}
