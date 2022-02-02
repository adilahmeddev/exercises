package methods

import (
	"bufio"
	"bytes"
	"exercises"
	"io"
	"io/fs"
	"strconv"
	"strings"
)

type NumberFile string

type Adder struct {
	history   map[int]int
	formatter Formatter
	numbers []int
}

func NewAdder(history map[int]int, formatter Formatter) *Adder {
	return &Adder{history: history, formatter: formatter}
}
func (a* Adder) LoadNumbers(args []string, fs fs.FS, numbers []int)  {


	typeOfArgs := a.ArgType(args)
	if typeOfArgs == "file" {
		fileNames := a.GetFiles(args)
		for _, fileName := range fileNames {
			file, err := fs.Open(string(fileName))
			if err != nil {
				panic(err)
			}
			a.InputFromFile(file)
		}

	}
	if typeOfArgs == "number" {
		a.InputFromArgs(args, numbers)
	}

}



func (a* Adder) ArgType(args []string) string {
	if args[0] == "--input-file" {
		return "file"
	}
	return "number"
}

func (a* Adder) GetFiles(args []string) []NumberFile {
	var numberFile []NumberFile
	if len(args)==0{
		numberFile = append(numberFile, NumberFile("input.txt"))
		return numberFile
	}

	for i := 0; i < len(args); i++ {
		if args[i] == "--input-file" {
			temp := NumberFile(args[i+1])
			numberFile = append(numberFile, temp)
			i++
		}
	}
	return numberFile
}

func (a* Adder) InputFromArgs(args []string, numbers []int) {
	for _, arg := range args {
		a.numbers = a.extractNumber(arg)
	}
}

func (a* Adder) InputFromFile(file fs.File) {
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
		a.numbers = a.readFromCSV(&buf)
	} else {
		a.numbers = a.readFromTxt(&buf)
	}

}

func (a* Adder) readFromTxt(file *bytes.Buffer) []int {
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = a.extractNumber(scanner.Text())
	}
	return numbers
}

func (a* Adder) readFromCSV(file *bytes.Buffer) []int {

	fileContents := make([]byte, file.Len())
	i, err := file.Read(fileContents)
	if err != nil {
		panic(strconv.Itoa(i) + " " + err.Error())
	}
	records := strings.Split(string(fileContents), ",")
	for _, number := range records {
		a.numbers = a.extractNumber(number)
	}
	return a.numbers
}

func (a* Adder) extractNumber(input string) []int {
	var numbers []int
	i, err := strconv.Atoi(input)
	if err == nil {
		if a.history[i] == 0 {
			numbers = append(a.numbers, i)
		}
		a.history[i]=1
	}
	return numbers
}

func (a* Adder) Format() string {
	return a.formatter.FormatNumber(exercises.Add(a.numbers...))
}

