package main

import (
	"exercises"
	"exercises/add/methods"
	"fmt"
	"net/http"
)


func MathServer(w http.ResponseWriter, r *http.Request) {
	nums := r.URL.Query()["num"]
	formatter := methods.Formatter{}
	adder := methods.NewAdder(map[int]int{}, formatter)
	numbers := []int{}
	numbers = adder.InputFromArgs(nums, numbers)
	fmt.Fprint(w, formatter.FormatNumber(exercises.Add(numbers...)) )
}


