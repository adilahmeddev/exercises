package numberLoader

import (
	"github.com/matryer/is"
	"os"
	"testing"
)

func TestNumberLoader(t *testing.T) {
	t.Run("load with no args", func(t *testing.T) {
		is := is.New(t)

		file, err := os.Create( "input.txt")
		if err != nil {
			is.NoErr(err)
		}
		file.Write([]byte(`5
10
7`))


		numberLoader := NewNumberLoader(os.DirFS("./"))
		want := []int{5,10,7}
		got, err := numberLoader.Load([]string{})


		is.NoErr(err)
		is.Equal(want, got)


		file.Close()
		err = os.Remove(file.Name())
		is.NoErr(err)

	})
	t.Run("load with input file flag", func(t *testing.T) {
		is := is.New(t)

		fileName := "input1.txt"
		file, err := os.Create(fileName)
		if err != nil {
			is.NoErr(err)
		}
		file.Write([]byte(`5
10
7`))

		fileName2 := "input2.txt"
		file2, err := os.Create( fileName2)
		if err != nil {
			is.NoErr(err)
		}
		file2.Write([]byte(`6
8
9`))



		numberLoader := NewNumberLoader(os.DirFS("./"))
		want := []int{5,10,7,6,8,9}
		got, err := numberLoader.Load([]string{"--input-file",fileName, "--input-file", fileName2})


		is.NoErr(err)
		is.Equal(want, got)


		file.Close()
		err = os.Remove(file.Name())
		is.NoErr(err)

		file2.Close()
		err = os.Remove(file2.Name())
		is.NoErr(err)

	})
}
