package main

import (
	"github.com/matryer/is"
	"testing"
	"testing/fstest"
)

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
		formattedNumber := FormatNumber(tc.number)
		is.Equal(formattedNumber, tc.want)
	}

}

func TestInsertIntoSlice(t *testing.T) {
	is := is.New(t)
	chars := []rune{'a', 'b', 'd'}
	expected := []rune{'a', 'b', 'c', 'd'}
	result := InsertIntoSlice(chars, 2, 'c')

	is.Equal(result, expected)
}

func TestInputFromFile(t *testing.T) {
	is := is.New(t)
	fs := fstest.MapFS{
		"input.txt" : &fstest.MapFile{
			Data:    []byte(`4
5
32
100
867543`)},
	}

	wantedNumbers := []int {4, 5, 32, 100, 867543}
	gotNumbers := []int{}
	gotNumbers = InputFromFile(gotNumbers, fs)

	is.Equal(gotNumbers, wantedNumbers)
}


