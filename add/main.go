package main

import (
	"bufio"
	"bytes"
	"exercises"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type NumberFile string


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
		typeOfArgs := ParseArgs(args)
		if typeOfArgs == "file"{
			fileNames := GetFiles(args)
			for _, fileName := range fileNames {
				file, err := fs.Open(string(fileName))
				if err != nil {
					panic(err)
				}
				numbers = InputFromFile(numbers, file)
			}

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

func GetFiles(args []string) []NumberFile {
	var numberFile []NumberFile

	for i := 0; i < len(args); i++ {
		if args[i] == "--input-file" {
			temp := NumberFile(args[i+1])
			numberFile = append(numberFile, temp)
			i++
		}
	}
	return numberFile
}

func InputFromArgs(args []string, numbers []int) []int {
	for _, arg := range args {
		numbers = extractNumber(arg, numbers)
	}
	return numbers
}

func InputFromFile(numbers []int, file fs.File) []int {
	defer file.Close()
	var buf bytes.Buffer
	tee := io.TeeReader(file, &buf)
	bReader := bufio.NewReaderSize(tee, 50)
	initialBytes := make([]byte, 50)
	initialBytes, _ = bReader.Peek(50)

	isCSV := false
	for _,char := range []rune(string(initialBytes)) {
		if char == ','{
			isCSV = true
		}
	}
	if isCSV{
		numbers = readFromCSV(numbers, &buf)
	} else {
		numbers =  readFromTxt(numbers, &buf)
	}


	return numbers
}

func readFromTxt(numbers []int, file *bytes.Buffer) []int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = extractNumber(scanner.Text(), numbers)
	}
	return numbers
}

func readFromCSV(numbers []int, file *bytes.Buffer) []int {

	fileContents := make([]byte,file.Len())
	i, err := file.Read(fileContents)
	if err != nil {

		panic(strconv.Itoa(i) + " " + err.Error())
	}
	records := strings.Split(string(fileContents), ",")
	for _, number := range records {
		numbers = extractNumber(number, numbers)
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
