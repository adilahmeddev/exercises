package main

import (
	"bytes"
	"exercises/add/methods"
	"github.com/matryer/is"
	"testing"
	"testing/fstest"
)


func TestRun(t *testing.T) {
	is := is.New(t)
	want:= "15"
	buffer := bytes.Buffer{}

	run([]string{"add","10", "5"}, &buffer)

	is.Equal(buffer.String(), want)
}

func TestFormatNumber(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		number int
		want   string
	}{
		{11111111,
			"11,111,111"},
		{999,
			"999"},
		{11500,
			"11,500"},
		{
			1234567890,
			"1,234,567,890"},
		{
			2523542352352355,
			"2,523,542,352,352,355"},
		{
			5555555555555555555,
			"5,555,555,555,555,555,555"},

	}

	for _, tc := range tests {
		formatter := methods.Formatter{}
		formattedNumber := formatter.FormatNumber(tc.number)
		is.Equal(formattedNumber, tc.want)
	}

}

func TestInsertIntoSlice(t *testing.T) {
	is := is.New(t)
	chars := []rune{'a', 'b', 'd'}
	expected := []rune{'a', 'b', 'c', 'd'}

	formatter := methods.Formatter{}
	result := formatter.InsertIntoSlice(chars, 2, 'c')

	is.Equal(result, expected)
}

func TestParseInput(t *testing.T){
	is := is.New(t)
	t.Run("from txt", func(t *testing.T) {
		fileInputArgs := []string{"--input-file", "input.txt"}
		numberInputArgs := []string{"20", "5", "2", "-3"}

		adder := methods.NewAdder(map[int]int{}, methods.Formatter{})

		isFile := adder.ArgType(fileInputArgs)
		isNumber := adder.ArgType(numberInputArgs)

		is.Equal(isFile, "file")
		is.Equal(isNumber, "number")
	})

	t.Run("from csv", func(t *testing.T) {
		fileInputArgs := []string{"--input-file", "input.csv"}
		numberInputArgs := []string{"20", "5", "2", "-3"}

		adder := methods.NewAdder(map[int]int{}, methods.Formatter{})

		isFile := adder.ArgType(fileInputArgs)
		isNumber := adder.ArgType(numberInputArgs)

		is.Equal(isFile, "file")
		is.Equal(isNumber, "number")
	})

}

func TestInputFromFile(t *testing.T) {
	t.Skip()
	is := is.New(t)
	fs := fstest.MapFS{
		"input.txt" : &fstest.MapFile{
			Data:    []byte(`4
5
32
100
867543`)},
	}

	file, err := fs.Open("input.txt")
	is.NoErr(err)
	wantedNumbers := []int {4, 5, 32, 100, 867543}
	gotNumbers := []int{}
	adder := methods.NewAdder(map[int]int{}, methods.Formatter{})
	adder.InputFromFile(file)

	is.Equal(gotNumbers, wantedNumbers)
}


