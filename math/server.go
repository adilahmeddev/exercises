package main

import (
	"exercises"
	"exercises/add/methods"
	"fmt"
	"net/http"
)


func MathServer(w http.ResponseWriter, r *http.Request) {
	nums := []string{}
	switch r.Header.Get("content-type"){
		case "application/x-www-form-urlencoded":
			r.ParseForm()
			nums = r.Form["num"]
	    case "application/json":

		default:
			nums = r.URL.Query()["num"]
	}
	formatter := methods.Formatter{}
	adder := methods.NewAdder(map[int]int{}, formatter)

	numbers := []int{}
	numbers = adder.InputFromArgs(nums, numbers)

	fmt.Fprint(w, formatter.FormatNumber(exercises.Add(numbers...)) )
}


