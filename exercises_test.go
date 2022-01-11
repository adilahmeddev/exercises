package exercises

import (
	"github.com/matryer/is"
	"testing"
)


func TestAdd(t *testing.T) {
	t.Run("can add two numbers together", func(t *testing.T) {
		is := is.New(t)
		wanted := 9

		got:= add(5,4)

		is.Equal(got, wanted)
	})
	t.Run("can add five numbers together", func(t *testing.T) {
		is := is.New(t)
		wanted := 15

		got:= add(5,4, 15, -4, -5)

		is.Equal(got, wanted)
	})

}
