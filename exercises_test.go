package exercises

import (
	"github.com/matryer/is"
	"testing"
)


func TestAdd(t *testing.T) {
	t.Run("can Add two numbers together", func(t *testing.T) {
		is := is.New(t)
		wanted := 9

		got:= Add(5,4)

		is.Equal(got, wanted)
	})
	t.Run("can Add five numbers together", func(t *testing.T) {
		is := is.New(t)
		wanted := 15

		got:= Add(5,4, 15, -4, -5)

		is.Equal(got, wanted)
	})

}
