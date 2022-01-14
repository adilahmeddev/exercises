package main

import (
	"exercises"
	"fmt"
	"os"
	"strconv"
)

func main(){
	args := os.Args[1:]
	var numbers []int

	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	sum := exercises.Add(numbers...)
	formattedNumber := FormatNumber(sum)
	fmt.Println(formattedNumber)

}

func FormatNumber(number int) string {
	commaCount := 0
	numberAsString:= strconv.Itoa(number)
	if number <= 9999 {
		return numberAsString
	} else {
		chars := []rune(numberAsString)
		for i := len(chars); i>=0; i--{
			if (len(chars)-i-commaCount)%3==0 && i != 0 && i != len(chars){
				chars = InsertIntoSlice(chars, i, ',')
				commaCount++
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
