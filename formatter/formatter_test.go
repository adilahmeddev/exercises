package formatter

import (
	"github.com/matryer/is"
	"testing"
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
		formatter := Formatter{}
		formattedNumber := formatter.FormatNumber(tc.number)
		is.Equal(formattedNumber, tc.want)
	}

}
