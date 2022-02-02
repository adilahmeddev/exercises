package main

import (
	"exercises/add/methods"
	"fmt"
	"io"
	"os"
)


func run(args []string, output io.Writer){
	var numbers []int
	fs := os.DirFS("add/")

	formatter := methods.Formatter{}
	adder := methods.NewAdder(map[int]int{}, formatter)

	adder.LoadNumbers(args, fs, numbers)


	formattedNumber := adder.Format()
	fmt.Fprint(output, formattedNumber)
}



func main(){
	args := os.Args[1:]
	run(args, os.Stdout)

}

