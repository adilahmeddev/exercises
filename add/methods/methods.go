package methods

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"strconv"
	"strings"
)

type NumberFile string

type Adder struct {
	history map[int]int
	formatter Formatter
}

func NewAdder(history map[int]int, formatter Formatter) *Adder {
	return &Adder{history: history, formatter: formatter}
}


func (a Adder) ParseArgs(args []string) string {
	if args[0] == "--input-file" {
		return "file"
	}
	return "number"
}

func (a Adder) GetFiles(args []string) []NumberFile {
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

func (a Adder) InputFromArgs(args []string, numbers []int) []int {
	for _, arg := range args {
		numbers = a.extractNumber(arg, numbers)
	}
	return numbers
}

func (a Adder) InputFromFile(numbers []int, file fs.File) []int {
	defer file.Close()
	var buf bytes.Buffer
	tee := io.TeeReader(file, &buf)
	bReader := bufio.NewReaderSize(tee, 50)
	initialBytes := make([]byte, 50)
	initialBytes, _ = bReader.Peek(50)

	isCSV := false
	for _, char := range []rune(string(initialBytes)) {
		if char == ',' {
			isCSV = true
		}
	}
	if isCSV {
		numbers = a.readFromCSV(numbers, &buf)
	} else {
		numbers = a.readFromTxt(numbers, &buf)
	}

	return numbers
}

func (a Adder) readFromTxt(numbers []int, file *bytes.Buffer) []int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = a.extractNumber(scanner.Text(), numbers)
	}
	return numbers
}

func (a Adder) readFromCSV(numbers []int, file *bytes.Buffer) []int {

	fileContents := make([]byte, file.Len())
	i, err := file.Read(fileContents)
	if err != nil {

		panic(strconv.Itoa(i) + " " + err.Error())
	}
	records := strings.Split(string(fileContents), ",")
	for _, number := range records {
		numbers = a.extractNumber(number, numbers)
	}
	return numbers
}

func (a Adder) extractNumber(input string, numbers []int) []int {
	i, err := strconv.Atoi(input)
	if err == nil {
		if a.history[i] == 0 {
			numbers = append(numbers, i)
		}
		a.history[i]=1
	}
	return numbers
}

type Formatter struct {

}
func (f Formatter) FormatNumber(number int) string {
	commaCount := 0
	numberAsString := strconv.Itoa(number)
	if number <= 9999 {
		return numberAsString
	} else {
		chars := []rune(numberAsString)
		for i := len(chars); i >= 0; i-- {
			if (len(chars)-i-commaCount)%3 == 0 && i != 0 && i != len(chars) {
				chars = f.InsertIntoSlice(chars, i, ',')
				commaCount++
			}
		}
		return string(chars)
	}
}

func (f Formatter) InsertIntoSlice(slice []rune, index int, char rune) []rune {
	slice = append(slice, 'x')
	copy(slice[index+1:], slice[index:])
	slice[index] = char
	return slice
}
