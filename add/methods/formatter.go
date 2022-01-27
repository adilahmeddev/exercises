package methods

import "strconv"

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
