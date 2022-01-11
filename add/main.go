package main

import (
	"bufio"
	"exercises"
	"fmt"
	"os"
	"strconv"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	var numbers []int

	fmt.Println("Enter all the numbers you want to add, and type \"stop\" when you're done")

	for input.Scan() {
		i, err := strconv.Atoi(input.Text())
		if err == nil {
			numbers = append(numbers, i)
		} else {
			if input.Text() == "stop" {
				break
			}
		}
	}
	sum := exercises.Add(numbers...)
	fmt.Println(sum)

}
