package main

import (
	"bufio"
	"exercises"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func main(){
	args := os.Args[1:]
	var numbers []int

	if len(args) == 0 {
		numbers = InputFromFile(numbers, os.DirFS("add/"))
	}
	numbers = InputFromArgs(args, numbers)

	sum := exercises.Add(numbers...)
	formattedNumber := FormatNumber(sum)
	fmt.Println(formattedNumber)
}

func InputFromArgs(args []string, numbers []int) []int {
	for _, arg := range args {
		numbers = extractNumber(arg, numbers)
	}
	return numbers
}

func InputFromFile(numbers []int, fs fs.FS) []int {
	file, err := fs.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = extractNumber(scanner.Text(), numbers)
	}
	return numbers
}

func extractNumber(input string, numbers []int) []int {
	i, err := strconv.Atoi(input)
	if err == nil {
		numbers = append(numbers, i)
	}
	return numbers
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
