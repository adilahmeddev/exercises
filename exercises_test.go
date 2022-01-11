package exercises

import (
	"testing"
	"github.com/matryer/is"
)


func TestAdd(t *testing.T) {
	is := is.New(t)
	wanted := 9

	got:= add(5,4)

	is.Equal(got, wanted)
}
