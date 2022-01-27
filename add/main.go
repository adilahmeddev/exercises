package main

import (
	"exercises"
	"exercises/add/methods"
	"fmt"
	"os"
)




func main(){
	fs := os.DirFS("add/")
	args := os.Args[1:]
	var numbers []int

	adder := methods.NewAdder(map[int]int{})
	if len(args) == 0 {
		file, err := fs.Open("input.txt")
		if err != nil {
			panic(err)
		}
		numbers = adder.InputFromFile(numbers,file)
	} else {
		typeOfArgs := adder.ParseArgs(args)
		if typeOfArgs == "file"{
			fileNames := adder.GetFiles(args)
			for _, fileName := range fileNames {
				file, err := fs.Open(string(fileName))
				if err != nil {
					panic(err)
				}
				numbers = adder.InputFromFile(numbers, file)
			}

		}
		if typeOfArgs == "number" {
			numbers = adder.InputFromArgs(args, numbers)
		}
	}

	sum := exercises.Add(numbers...)
	formattedNumber := adder.FormatNumber(sum)
	fmt.Println(formattedNumber)
}

