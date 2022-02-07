package main

import (
	"calculator/cmd/add/numberLoader"
	"calculator/formatter"
	sum2 "calculator/sum"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)


func MathServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer " +os.Getenv("super_secret_api_key"){
	//if r.Header.Get("Authorization") != "Bearer SUPER_SECRET_API_KEY"{
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			w.WriteHeader(http.StatusForbidden)
			return
		}

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
	loader := numberLoader.NewNumberLoader(os.DirFS("./"))

	numbers, err := loader.Load(nums)
	if err != nil {
		fmt.Fprint(w,err)
		return
	}
	sum := sum2.Sum(numbers...)

	formatter := formatter.Formatter{}
	fmt.Fprint(w, formatter.FormatNumber(sum))
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