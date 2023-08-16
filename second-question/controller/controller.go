package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Response struct {
	Numbers []int `json:"numbers"`
}

func GetNumbers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get numbers - ----")
	queryValues := r.URL.Query()

	fmt.Println("queryValues: ", queryValues)

	urls, err := queryValues["url"]
	if err != true {
		fmt.Println("Error: ", err)
	}
	fmt.Println("numbers: ", urls)

	NumbersSet := make(map[int]bool)

	for _, url := range urls {
		fmt.Println("url: ", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		var responseObj Response
		err = json.Unmarshal(body, &responseObj)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("responseObj: ", responseObj.Numbers)

		for _, num := range responseObj.Numbers {
			NumbersSet[num] = true
		}
	}

	fmt.Println("-----------------------")

	var keys []int
	for k := range NumbersSet {
		keys = append(keys, k)
	}

	fmt.Println("Unsorted keys: ", keys)
	sort.Ints(keys)
	fmt.Println("Sorted keys: ", keys)

	response := Response{Numbers: keys}
	responseJSON, _ := json.MarshalIndent(response, "", "")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
