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

type numbers struct {
	nums []int `json:"nums"`
}
func MathServer(w http.ResponseWriter, r *http.Request) {
	nums := []string{}
	defer r.Body.Close()
	switch r.Header.Get("content-type"){
		case "application/x-www-form-urlencoded":
			r.ParseForm()
			nums = r.Form["num"]
	    case "application/json":
			var result map[string][]int
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			if err = json.Unmarshal(bytes, &result); err != nil {
				panic(err)
			}
			fmt.Println(result)
			for _, num := range result["nums"]{
				nums = append(nums, strconv.Itoa(num))
			}
			fmt.Println(nums)
	default:
			nums = r.URL.Query()["num"]
	}
	formatter := methods.Formatter{}
	adder := methods.NewAdder(map[int]int{}, formatter)

	numbers := []int{}
	numbers = adder.InputFromArgs(nums, numbers)

	fmt.Fprint(w, formatter.FormatNumber(exercises.Add(numbers...)) )
}


