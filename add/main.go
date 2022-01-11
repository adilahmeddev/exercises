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
	formattedNumber := FormatNumber(sum)
	fmt.Println(formattedNumber)

}

func FormatNumber(number int) string {
	numberAsString:= strconv.Itoa(number)
	if number <= 9999 {
		return numberAsString
	} else {
		chars := []rune(numberAsString)
		for i := len(chars); i>=0; i--{
			if (i+1)%3==0 && i != 0 && i != len(chars){
				chars = InsertIntoSlice(chars, i, ',')
			}
		}
		return string(chars)
	}
}

func InsertIntoSlice(slice []rune, index int, char rune) []rune {
	slice = append(slice, 'x')
	copy(slice[index+1:], slice[index:])
	slice[index] = char
	return slice
}
