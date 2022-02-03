package sum

import (
	"github.com/matryer/is"
	"testing"
)

func TestSum(t *testing.T) {
	is := is.New(t)
	t.Run("adding two numbers", func(t *testing.T) {
		want := 15
		got := Sum(5,10)

		is.Equal(got, want)
	})
	t.Run("adding three numbers", func(t *testing.T) {
		want := 22
		got := Sum(5,10, 7)

		is.Equal(got, want)
	})

}
