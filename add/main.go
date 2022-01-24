package main

import (
	"bufio"
	"encoding/csv"
	"exercises"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func main(){
	fs := os.DirFS("add/")
	args := os.Args[1:]
	var numbers []int

	if len(args) == 0 {
		file, err := fs.Open("input.txt")
		if err != nil {
			panic(err)
		}
		numbers = InputFromFile(numbers,file)
	} else {
		fmt.Println(args[1])
		typeOfArgs := ParseArgs(args)
		if typeOfArgs == "file"{
			file, err := fs.Open(args[1])
			if err != nil {
				panic(err)
			}
			numbers = InputFromFile(numbers, file)
		}
		if typeOfArgs == "number" {
			numbers = InputFromArgs(args, numbers)
		}
	}

	sum := exercises.Add(numbers...)
	formattedNumber := FormatNumber(sum)
	fmt.Println(formattedNumber)
}

func ParseArgs(args []string) string{
	if args[0] == "--input-file" {
		return "file"
	}
	return "number"
}

func InputFromArgs(args []string, numbers []int) []int {
	for _, arg := range args {
		numbers = extractNumber(arg, numbers)
	}
	return numbers
}

func InputFromFile(numbers []int, file fs.File) []int {
	defer file.Close()
	fileInfo, _ := file.Stat()
	if extension := strings.Split(fileInfo.Name(), "."); extension[1] == "csv"{
		numbers = readFromCSV(numbers, file)
	}
	if extension := strings.Split(fileInfo.Name(), "."); extension[1] == "txt" {
		numbers = readFromTxt(numbers, file)
	}

	return numbers
}

func readFromTxt(numbers []int, file fs.File) []int {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = extractNumber(scanner.Text(), numbers)
	}
	return numbers
}

func readFromCSV(numbers []int, file fs.File) []int {
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	numberRecords := records[0]
	for _, number := range numberRecords {
		numbers = extractNumber(string(number), numbers)
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
