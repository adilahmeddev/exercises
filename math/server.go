package main

import (
	"encoding/json"
	"exercises"
	"exercises/add/methods"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)


func MathServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer SUPER_SECRET_API_KEY"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	nums := []string{}
	switch r.Header.Get("content-type"){
		case "application/x-www-form-urlencoded":
			r.ParseForm()
			nums = r.Form["num"]
	    case "application/json":
			nums = parseJSON(r, nums)
		default:
			nums = r.URL.Query()["num"]
	}
	formatter := methods.Formatter{}
	adder := methods.NewAdder(map[int]int{}, formatter)

	numbers := []int{}
	numbers = adder.InputFromArgs(nums, numbers)

	fmt.Fprint(w, formatter.FormatNumber(exercises.Add(numbers...)))
}

func parseJSON(r *http.Request, nums []string) []string {
	var result map[string][]int
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &result); err != nil {
		panic(err)
	}
	for _, num := range result["nums"] {
		nums = append(nums, strconv.Itoa(num))
	}
	return nums
}


