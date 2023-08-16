package controller

import (
	"fmt"
	"net/http"
)

func GetNumbers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get numbers - ----")
	queryValues := r.URL.Query()

	fmt.Println("queryValues: ", queryValues)

	urls, err := queryValues["url"]
	if err != true {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("numbers: ", urls)

}
